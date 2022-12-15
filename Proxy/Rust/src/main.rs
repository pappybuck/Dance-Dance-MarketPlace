use axum::{
    extract::{Extension, Host},
    http::{uri::Uri, Request},
    routing::{any},
    Router, body::Bytes, response::{AppendHeaders, Response, IntoResponse},
};
use hyper::{client::HttpConnector, Body, header::SET_COOKIE};
use tower::{ServiceExt, ServiceBuilder};
use tower_http::{
    add_extension::AddExtensionLayer,
    compression::CompressionLayer,
    trace::{TraceLayer, DefaultMakeSpan}
};
use std::{net::SocketAddr, sync::Arc, time::Duration, collections::HashMap};
type Client = hyper::client::Client<HttpConnector, Body>;

struct State {
    client: Client,
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    let state = State {
        client: Client::new(),
    };

    let base_handler = Router::new().route("/*path", any(handler));

    let middleware = ServiceBuilder::new()
        .layer(AddExtensionLayer::new(Arc::new(state)))
        .layer(
            TraceLayer::new_for_http()
            .on_body_chunk(|chunk: &Bytes, latency: Duration, _: &tracing::Span| {
                tracing::trace!(size_bytes = chunk.len(), latency = ?latency, "sending body chunk")
            })
            .make_span_with(DefaultMakeSpan::new())
        )
        .layer(CompressionLayer::new())
        .into_inner();

    let app = Router::new()
        .route("/*path", any(
            |Host(hostname): Host, request: Request<Body>| async move {
                match hostname.as_str(){
                    _ => base_handler.oneshot(request).await,
                }
            }
        ))
        .layer(middleware);
    
    let addr = SocketAddr::from(([0, 0, 0, 0], 4000));
    // println!("reverse proxy listening on {}", addr);
    tracing::info!("reverse proxy listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}

async fn handler(
    Extension(state): Extension<Arc<State>>,
    mut req: Request<Body>,
) -> Response {
    let host = req.headers().get("host").unwrap().to_str().unwrap();
    tracing::info!("host: {}", host);
    let path = req.uri().path();
    let path_query = req
        .uri()
        .path_and_query()
        .map(|v| v.as_str())
        .unwrap_or(path);
    let mut paths = HashMap::new();
    paths.insert("redis.localhost".to_string(), "http://redis:8001".to_string());
    paths.insert("auth.localhost".to_string(), "http://auth:80".to_string());
    if paths.contains_key(host) {
        let uri = format!("{}{}", paths.get(host).unwrap(), path_query);

        *req.uri_mut() = Uri::try_from(uri).unwrap();
    
        let res = state.client.request(req).await;
        if res.is_err() {
            tracing::error!("error: {}", res.err().unwrap());
            return Response::builder()
                .status(500)
                .body(Body::empty())
                .unwrap().into_response();
        } else {

        }
    } else {
        tracing::error!("host not found: {}", host);
        return Response::builder()
            .status(404)
            .body(Body::empty())
            .unwrap().into_response();
    }
}



async fn BuildResponse(mut res : Body) -> impl IntoResponse {
    (
        AppendHeaders([(SET_COOKIE, "foo=bar")]),
        res,
    )
}
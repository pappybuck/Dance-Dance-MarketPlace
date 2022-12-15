import { GraphQLClient } from "graphql-request";
import {createClient } from "urql";

export const client = createClient({
    url: "https://api.patrickbuck.net/graphql/query",
    // url: "http://api.localhost/graphql/query",
    // url: "http://localhost:8080/query",
});
//"http://graphql.localhost:80/query"
// export const ssrClient = new GraphQLClient("http://graphql:4000/graphql/query");
export const ssrClient = new GraphQLClient(process.env.SSR_GRAPHQL_URL!);
// export const ssrClient = new GraphQLClient("http://localhost:4000/graphql/query");
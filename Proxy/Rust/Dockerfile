# Rust as the base image
FROM rust:1.65.0 as build

# Create a new empty shell project
RUN USER=root cargo new --bin proxy
WORKDIR /proxy

# Copy our manifests
COPY ./Cargo.lock ./Cargo.lock
COPY ./Cargo.toml ./Cargo.toml

# Build only the dependencies to cache them
RUN cargo build --release
RUN rm src/*.rs

# Copy the source code
COPY ./src ./src

# Build for release.
RUN rm ./target/release/deps/Proxy*
RUN cargo build --release

# The final base image
FROM debian:buster-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy from the previous build
COPY --from=build /proxy/target/release/Proxy /usr/src/Proxy
# COPY --from=build /holodeck/target/release/holodeck/target/x86_64-unknown-linux-musl/release/holodeck .
EXPOSE 4000
# Run the binary
CMD ["/usr/src/Proxy"]
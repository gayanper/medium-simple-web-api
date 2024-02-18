FROM --platform=linux/arm64 ghcr.io/graalvm/graalvm-community:21 AS builder

WORKDIR /build

COPY src ./src
COPY .mvn ./.mvn
COPY pom.xml .
COPY mvnw .

RUN ./mvnw -Pnative native:compile

FROM ubuntu:24.04

ENV DB_HOST=
ENV ENABLE_VT=

COPY --from=builder /build/target/employee-portal-sb /employee-portal-sb
ENTRYPOINT [ "/employee-portal-sb" ]
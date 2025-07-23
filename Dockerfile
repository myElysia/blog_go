FROM alpine
LABEL authors="elysia"
WORKDIR /build
COPY blogGo .
CMD ["./blogGo"]
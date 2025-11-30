load("@rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/gaespinoza/snake

load("@gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

go_library(
    name = "snake_lib",
    srcs = ["main.go"],
    importpath = "github.com/gaespinoza/snake",
    visibility = ["//visibility:private"],
    deps = [
        "//models",
        "//state",
        "@org_gioui//app",
        "@org_gioui//font/gofont",
        "@org_gioui//io/key",
        "@org_gioui//op",
        "@org_gioui//text",
        "@org_gioui//widget/material",
    ],
)

go_binary(
    name = "snake",
    embed = [":snake_lib"],
    visibility = ["//visibility:public"],
)

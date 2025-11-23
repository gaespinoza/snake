load("@rules_go//go:def.bzl", "go_binary", "go_library")

# gazelle:prefix github.com/gaespinoza/snake

load("@gazelle//:def.bzl", "gazelle")

gazelle(name = "gazelle")

go_library(
    name = "snake_lib",
    srcs = ["main.go"],
    importpath = "github.com/gaespinoza/snake",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "snake",
    embed = [":snake_lib"],
    visibility = ["//visibility:public"],
)

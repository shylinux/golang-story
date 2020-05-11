Volcanos("onengine", { river: {
    "main": {name: "main", storm: {
        "main": {name: "main", action: [
            {name: "main", help: "main", inputs: [
                {type: "text", name: "path", value: "src/main.shy"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], group: "web.wiki", index: "word", feature: {
                display: "local/wiki/word",
            }},
        ]},
        "project": {name: "project", action: [
            {name: "main", help: "main", inputs: [
                {type: "text", name: "path", value: "src/project/project.shy"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], group: "web.wiki", index: "word", feature: {
                display: "local/wiki/word",
            }},
            {name: "project", help: "project", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "执行", action: "auto"},
            ], group: "web.code.project", index: "project", feature: {
                display: "/publish/project",
            }},
        ]},
        "compile": {name: "compile", action: [
            {name: "main", help: "main", inputs: [
                {type: "text", name: "path", value: "src/compile/compile.shy"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], group: "web.wiki", index: "word", feature: {
                display: "local/wiki/word",
            }},
            {name: "compile", help: "compile", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "执行", action: "auto"},
            ], group: "web.code.compile", index: "compile", feature: {
                display: "/publish/compile",
            }},
        ]},
        "runtime": {name: "runtime", action: [
            {name: "main", help: "main", inputs: [
                {type: "text", name: "path", value: "src/runtime/runtime.shy"},
                {type: "button", name: "查看", action: "auto"},
                {type: "button", name: "返回"},
            ], group: "web.wiki", index: "word", feature: {
                display: "local/wiki/word",
            }},
            {name: "runtime", help: "runtime", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "执行", action: "auto"},
            ], group: "web.code.runtime", index: "runtime", feature: {
                display: "/publish/runtime",
            }},
        ]},
        "hello": {name: "hello", action: [
            {name: "hello", help: "hello", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "执行", action: "auto"},
            ], engine: function(event, can, msg, pane, cmds, cb) {
                msg.Echo("hello world")
                typeof cb == "function" && cb(msg)
            }},
        ]},
        "world": {name: "world", action: [
            {name: "world", help: "world", inputs: [
                {type: "text", name: "one", value: "pwd"},
                {type: "button", name: "执行", action: "auto"},
            ], group: "cli", index: "system"},
        ]},
    }},
}, })

Volcanos(chat.ONIMPORT, {help: "导入数据", _init: function(can, msg, cb, target) {
		can.base.isFunc(cb) && cb(msg)
	},
	_show2: function(can, cb) {
		can.onimport._plugin(can, function(sub) {
			can.base.isFunc(cb) && cb(function(tree) {
				sub.Action(ice.VIEW, "纵向"), sub._tree = tree, sub.onimport.layout(sub)
				can.page.Select(can, sub._target, html.DIV_TOGGLE, function(target) { can.onmotion.hidden(can, target) })
			}, sub)
		}, ["/plugin/story/spide.js"])
	},
	_plugin: function(can, cb, list) {
		can.onappend.plugin(can, {index: "can.plugin"}, function(sub) {
			can.onmotion.hidden(can, sub._legend)
			can.onmotion.hidden(can, sub._option)
			can.onmotion.hidden(can, sub._action)
			can.onmotion.hidden(can, sub._status)
			can.onappend._status(sub, ["name", "count", "compare", "swap"])
			can.page.style(can, sub._target, html.MARGIN, 0, html.FLOAT, html.LEFT)
			sub.ConfHeight(can.onexport.height(can)+html.ACTION_HEIGHT), sub.ConfWidth(can.onexport.width(can)-2*html.PLUGIN_MARGIN)
			can.page.style(can, sub._output, html.HEIGHT, can.onexport.height(can)-html.ACTION_HEIGHT+html.PLUGIN_MARGIN, html.WIDTH, sub.ConfWidth())
			sub._display_output = function(table, msg) { sub.sup = can, delete(sub.onaction.list)
				sub.require((list||[]).concat(["/plugin/local/wiki/draw.js", "/plugin/local/wiki/draw/path.js", "/plugin/table.js"]), function() {
					sub.onappend._action(sub, sub.onaction.list, sub._action), can.base.isFunc(cb) && cb(sub)
					can.page.Select(can, sub._target, html.DIV_TOGGLE, function(target) { can.onmotion.hidden(can, target) })
				})
			}
		}, can._output)
	},
})
Volcanos(chat.ONACTION, {help: "操作数据", list: [
		"全部", "排序", ["算法", "红黑树"],
		[html.HEIGHT, ice.AUTO, "100", "200", "300", "500", "1", "2", "3", ice.AUTO],
		[html.WIDTH, "1", "2", "3"],
		[html.SPEED, "50", "100", "300", "500"],
		[cli.COLOR, ice.AUTO, cli.BLACK, cli.WHITE, cli.RED],
		[chat.TITLE, html.OUTPUT, html.STATUS],
	],
	"排序": function(event, can) {
		var msg = can.request(event)
		msg.Push("data", "12")
		msg.Push("data", "23")
		msg.Push("data", "34")
		msg.Push("data", "40")
		msg.Push("data", "45")
		msg.Push("data", "67")
		msg.Push("data", "78")
		msg.Push("data", "89")
		msg.Push("data", "90")
		msg.Push("data", "100")
		msg.Push("data", "110")
		msg.Push("data", "120")
		msg.Push("data", "130")
		msg.Push("data", "140")
		var data = can._msg.Table(), stat = {comp: 0, swap: 0}, record = []
		can.onimport._show2(can, function(show, sub) {
			can.onfigure[can.Action("算法")](event, can, data, record, stat, function(tree) {
				show(tree), can.onexport.title(can, sub, data, stat)
			})
		})
	},
})
Volcanos(chat.ONFIGURE, {help: "数据算法",
	"红黑树": function(event, can, data, record, stat, show) {
		function colors(n, color) { return n.meta.color = color }
		function setroot(tree, n, g) { colors(n, cli.BLACK), colors(g, cli.RED)
			if (!g.back) { tree[""] = n } else { g.back.list[g.back.list[0] == g? 0: 1] = n } n.back = g.back
		}
		function lrotate(g, n) { g.list[1] = n.list[0], n.list[0].back = g, n.list[0] = g, g.back = n }
		function rrotate(g, n) { g.list[0] = n.list[1], n.list[1].back = g, n.list[1] = g, g.back = n }
		function rotate(tree, n) { colors(n, n.back? cli.RED: cli.BLACK) 

			while (n.back && n.back.back && n.back.meta.color == cli.RED) { var p = n.back, g = n.back.back
				if (g.list[0].meta && g.list[0].meta.color == cli.RED && g.list[1].meta && g.list[1].meta.color == cli.RED) {
					colors(g.list[0], cli.BLACK), colors(g.list[1], cli.BLACK), colors(g, cli.RED)
					if (!g.back) { colors(g, cli.BLACK); break }
					n = g
					continue
				}
/*
 *  LL:  g  =>  p         LR: g   =>   g  =>  n         RL: g  =>  g   =>   n        RR: g    =>    p     
 *      / \    / \           / \      / \    / \           / \    / \      / \          / \        / \    
 *     p      n   g         p        n      p   g             p      n    g   p            p      g   n   
 *    / \        / \       / \      /          / \           / \      \  /                / \    /        
 *   n                        n    p                        n          p                     n            
 */ 
				if (g.list[0] == p) {
					if (p.list[0] == n) { // LL
						setroot(tree, p, g), rrotate(g, p)
					} else { // LR
						setroot(tree, n, g), lrotate(p, n), rrotate(g, n)
					}
				} else {
					if (p.list[0] == n) { // RL
						setroot(tree, n, g), rrotate(p, n), lrotate(g, n)
					} else { // RR
						setroot(tree, p, g), lrotate(g, p)
					}
				}
			}
		}

		var tree = {}; can.core.Next(data, function(item, next) { var node = {name: item.data, meta: {}, list: [{name: " "}, {name: " "}]}
			var root = tree[""]; if (!root) { root = tree[""] = node } else { while (root) {
				if (can.onaction.comp(can, [{data: item.data}, {data: root.name}], 0, ">", 1, stat)) {
					if (!root.list[1].meta) { node.back = root, root.list[1] = node; break } root = root.list[1]
				} else {
					if (!root.list[0].meta) { node.back = root, root.list[0] = node; break } root = root.list[0]
				}
			} } rotate(tree, node), show(tree)
			// can.onaction.next(can, next)
			can.__next = next, can.user.toastSuccess(can, node.name)
		}, function() { can.user.toastSuccess(can) })
	},
	"红黑树": function(event, can, data, record, stat, show) { var tree = {}
		can.core.Next(data, function(item, next) { var fork = {name: item.data, meta: {color: cli.RED}, list: [{}, {}]}
			var node = tree[""]; if (!node) { node = tree[""] = fork, fork.meta.color = cli.BLACK } else { while (node) {
				if (can.onaction.comp(can, [{data: item.data}, {data: node.name}], 0, ">", 1, stat)) {
					if (!node.list[1].name) {
						fork.back = node
						node.list[1] = fork;
						break
					} node = node.list[1]
				} else {
					if (!node.list[0].name) {
						fork.back = node
						node.list[0] = fork;
						break
					} node = node.list[0]
				}
			} } show(tree), can.onaction.next(can, next)
		}, function() { can.user.toastSuccess(can) })
	},
})
Volcanos(chat.ONEXPORT, {help: "导出数据",
	height: function(can) {
		return parseInt(can.Action(html.HEIGHT) == ice.AUTO? can.ConfHeight(): parseInt(can.Action(html.HEIGHT)) < 10? can.ConfHeight()/can.Action(html.HEIGHT) :can.Action(html.HEIGHT))
	},
	width: function(can) {
		return can.ConfWidth()/can.Action(html.WIDTH)
	},
	color: function(can) {
		if (can.Action(cli.COLOR) == ice.AUTO) {
			switch (can.getHeader(chat.TOPIC)) {
				case cli.WHITE: return cli.BLACK
				case cli.BLACK: return cli.WHITE
			}
		}
		return can.Action(cli.COLOR)
	},
	title: function(can, sub, data, stat) {
		if (can.Action(chat.TITLE) == html.STATUS) { can.onmotion.toggle(can, sub._status, true)
			sub.Status({name: can.Action("算法"), count: data.length, compare: stat.comp, swap: stat.swap})
		} else {
			sub.onimport.draw({}, sub, {shape: svg.TEXT, style: {
				inner: can.Action("算法")+ice.SP+can.base.joinKV([mdb.COUNT, data.length, "compare", stat.comp, "swap", stat.swap]),
				fill: can.onexport.color(can), "text-anchor": "start", "dominant-baseline": "text-before-edge",
			}, point: [{x: 0, y: 0}]})
		}
	},
})

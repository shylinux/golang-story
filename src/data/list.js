Volcanos(chat.ONIMPORT, {help: "导入数据", _init: function(can, msg, cb, target) {
		can.onmotion.clear(can), can.base.isFunc(cb) && cb(msg)
		can.Option(mdb.ZONE) == "" && can.onappend.table(can, msg)
	},
	_show: function(can, cb) {
		can.onimport._plugin(can, function(sub) { var msg = can._msg, list = []
			var min, max, step = (can.onexport.width(can)-2*html.PLUGIN_MARGIN)/msg.Length(), height = sub.ConfHeight(can.onexport.height(can)+html.ACTION_HEIGHT)-html.ACTION_HEIGHT-2*html.PLUGIN_MARGIN
			msg.Table(function(value, index) { (index == 0 || value.data < min) && (min = parseInt(value.data)), (index == 0 || value.data > max) && (max = parseInt(value.data)) })
			sub.onimport._show(sub, sub.request()), msg.Table(function(value, index) { list.push(sub.onimport.draw({}, sub, {
				shape: svg.RECT, style: {fill: can.onexport.color(can)}, point: [
					{x: index*step, y: height}, {x: index*step+step, y: height - (value.data-min) / (max-min) * height}
				]
			})) }), can.base.isFunc(cb) && cb(list, sub)
		})
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
Volcanos(chat.ONACTION, {help: "操作数据", list: [mdb.INSERT, "数据",
		"全部", "二叉树", "排序", ["算法", "红黑树", "快速排序", "选择排序", "交换排序", "插入排序"],
		"下一步",
		[html.HEIGHT, ice.AUTO, "100", "200", "300", "500", "1", "2", "3", ice.AUTO],
		[html.WIDTH, "1", "2", "3"],
		[html.SPEED, "50", "100", "300", "500"],
		[cli.COLOR, ice.AUTO, cli.BLACK, cli.WHITE, cli.RED],
		[chat.TITLE, html.OUTPUT, html.STATUS],
	],
	comp: function(can, data, i, op, j, stat) { stat.comp++
		switch (op) {
			case "<=": return parseInt(data[i].data) <= parseInt(data[j].data)
			default: return parseInt(data[i].data) > parseInt(data[j].data)
		}
	},
	swap: function(can, data, i, j, record, stat) { stat.swap++
		var tmp = data[i]; data[i] = data[j], data[j] = tmp
		record.push({action: "fill", index: i, value: cli.RED})
		record.push({action: "fill", index: j, value: cli.GREEN})
		record.push({action: "swap", index: i, value: j})
		record.push({action: "fill", index: i, value: can.onexport.color(can)})
		record.push({action: "fill", index: j, value: can.onexport.color(can)})
	},
	next: function(can, next) {
		can.core.Timer(parseInt(can.Action(html.SPEED)), next)
	},

	"全部": function(event, can) { can.Action(html.HEIGHT, "2"), can.Action(html.WIDTH, "2")
		can.core.Next(["选择排序", "交换排序", "快速排序", "插入排序"], function(method, next) {
			can.Action("算法", method), can.onaction["排序"](event, can)
			can.onmotion.delay(can, next, 100)
		})
	},
	"二叉树": function(event, can) { can.Action(html.HEIGHT, "2"), can.Action(html.WIDTH, "1")
		can.core.Next(["插入排序", "红黑树"], function(method, next) {
			can.Action("算法", method), can.onaction["排序"](event, can)
			can.onmotion.delay(can, next, 100)
		})
	},
	"数据": function(event, can) { can.onappend.table(can, can._msg)
		can.page.Select(can, can._output, html.TABLE_CONTENT, function(target) {
			can.page.style(can, target, html.MAX_HEIGHT, "300", html.DISPLAY, html.BLOCK)
		})
	},
	"下一步": function(event, can) {
		can.__next()
	},
	"排序": function(event, can) { var data = can._msg.Table(), stat = {comp: 0, swap: 0}, record = []
		if (can.Action("算法") == "插入排序" || can.Action("算法") == "红黑树") {
			can.onimport._show2(can, function(show, sub) {
				can.onfigure[can.Action("算法")](event, can, data, record, stat, function(tree) {
					show(tree), can.onexport.title(can, sub, data, stat)
				})
			})
			return
		}
		can.onimport._show(can, function(list, sub) { can.onfigure[can.Action("算法")](event, can, data, record, stat)
			can.onexport.title(can, sub, data, stat), can.onmotion.delay(can, function() { var count = 0
				can.core.Next(record, function(step, next) {
					switch (step.action) {
						case "fill":
							list[step.index].Value("fill", step.value)
							next()
							break
						case "swap": count++
							var x = list[step.index].Val("x")
							list[step.index].Val("x", list[step.value].Val("x"))
							list[step.value].Val("x", x)
							var node = list[step.index]
							list[step.index] = list[step.value]
							list[step.value] = node
							can.onaction.next(can, next)
							break
					}
				}, function() { can.user.toastSuccess(can) })
			})
		})
	},
})
Volcanos(chat.ONFIGURE, {help: "数据算法",
	"快速排序": function(event, can, data, record, stat) {
		function qsort(data, left, right) { var i = left, j = right
			while (i < j) {
				while (i < j && can.onaction.comp(can, data, i, "<=", j, stat)) { j-- } if (i >= j) { break }
				can.onaction.swap(can, data, i, j, record, stat)
				while (i < j && can.onaction.comp(can, data, i, "<=", j, stat)) { i++ } if (i >= j) { break }
				can.onaction.swap(can, data, i, j, record, stat)
			}
			left < i-1 && qsort(data, left, i-1), i+1 < right && qsort(data, i+1, right)
		} qsort(data, 0, data.length-1)
	},
	"选择排序": function(event, can, data, record, stat) {
		for (var i = 0; i < data.length-1; i++) {
			for (var j = i+1; j < data.length; j++) {
				if (can.onaction.comp(can, data, i, ">", j, stat)) {
					can.onaction.swap(can, data, i, j, record, stat)
				}
			}
		}
	},
	"交换排序": function(event, can, data, record, stat) {
		for (var i = 1; i < data.length; i++) {
			for (var j = 0; j < data.length-i; j++) {
				if (can.onaction.comp(can, data, j, ">", j+1, stat)) {
					can.onaction.swap(can, data, j, j+1, record, stat)
				}
			}
		}
	},
	"插入排序": function(event, can, data, record, stat, show) {
		var tree = {}; can.core.Next(data, function(item, next) { var node = {name: item.data, meta: {color: cli.RED}, list: [{}, {}]}
			var root = tree[""]; if (!root) { root = tree[""] = node } else { while (root) {
				if (can.onaction.comp(can, [{data: item.data}, {data: root.name}], 0, ">", 1, stat)) {
					if (!root.list[1].meta) { root.list[1] = node; break } root = root.list[1]
				} else {
					if (!root.list[0].meta) { root.list[0] = node; break } root = root.list[0]
				}
			} } show(tree), can.onaction.next(can, next), node.meta.color = can.onexport.color(can)
		}, function() { can.user.toastSuccess(can) })
	},
	"红黑树": function(event, can, data, record, stat, show) {
		function colors(n, color) { return n.meta.color = color }
		function isred(p) { return p && p.meta && p.meta.color == cli.RED }
		function setroot(tree, n, g) { colors(n, cli.BLACK), colors(g, cli.RED)
			if (!g.back) { tree[""] = n } else { g.back.list[g.back.list[0] == g? 0: 1] = n } n.back = g.back
		}
		function isleft(p, g) { return p == g.list[0] }
		function lrotate(p, g) { g.list[1] = p.list[0], p.list[0].back = g, p.list[0] = g, g.back = p, stat.swap++ }
		function rrotate(p, g) { g.list[0] = p.list[1], p.list[1].back = g, p.list[1] = g, g.back = p, stat.swap++ }
		function rotate(tree, n) { colors(n, n.back? cli.RED: cli.BLACK) 
			while (isred(n.back) && n.back.back) { var p = n.back, g = n.back.back
				if (isred(g.list[0]) && isred(g.list[1])) {
					colors(g.list[0], cli.BLACK), colors(g.list[1], cli.BLACK), colors(g, cli.RED)
					if (!g.back) { colors(g, cli.BLACK); break } else { n = g; continue }
				}
				/*
				 * 根节点是黑色
				 * 节点是红色或黑色
				 * 红色节点不能相邻
				 * 新节点插入前是红色，随后进行调整
				 */
/*
 *  LL:  g  =>  p        LR: g   =>   g  =>  n        RL: g  =>  g   =>   n        RR: g    =>    p    
 *      / \    / \          / \      / \    / \          / \    / \      / \          / \        / \   
 *     p      n   g        p        n      p   g            p      n    g   p            p      g   n  
 *    / \        / \      / \      /          / \          / \      \  /                / \    /       
 *   n                       n    p                       n          p                     n           
 */ 
				if (isleft(p, g)) {
					if (isleft(n, p)) { // LL
						setroot(tree, p, g), rrotate(p, g)
					} else { // LR
						setroot(tree, n, g), lrotate(n, p), rrotate(n, g)
					}
				} else {
					if (isleft(n, p)) { // RL
						setroot(tree, n, g), rrotate(n, p), lrotate(n, g)
					} else { // RR
						setroot(tree, p, g), lrotate(p, g)
					}
				}
				break
			}
		}

		var tree = {}; can.core.Next(data, function(item, next) { var node = {name: item.data, meta: {}, list: [{name: " "}, {name: " "}]}
			var root = tree[""]; if (!root) { root = tree[""] = node } else { while (root) {
				if (can.onaction.comp(can, [{data: item.data}, {data: root.name}], 0, ">", 1, stat)) {
					if (!root.list[1].meta) { node.back = root, root.list[1] = node; break } root = root.list[1]
				} else {
					if (!root.list[0].meta) { node.back = root, root.list[0] = node; break } root = root.list[0]
				}
			} } rotate(tree, node), show(tree), can.onaction.next(can, next) // can.__next = next, can.user.toastSuccess(can, node.name)
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

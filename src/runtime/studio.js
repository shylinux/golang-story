Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) { can.onmotion.hidden(can, can._status)
		can.ui = can.onappend.layout(can), can.onimport._project(can, msg)
	},
	_project: function(can, msg) {
		can.onimport.tree(can, msg.Table(function(value) { value.nick = value.path
			value._label = [{text: [value.version? `(${value.version})`: "", "", nfs.VERSION]}]
			return value
		}), function(event, value, target) { if (value._msg) { return can.onimport._content(can, value._msg, value, target) }
			can.run(event, [value.path], function(msg) { value._msg = msg, can.onimport._content(can, value._msg, value, target) })
		})
	},
	_content: function(can, msg, value, target) { can.onimport.tabsCache(can, value, target, msg) },
})
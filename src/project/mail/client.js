Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) { can.onmotion.hidden(can, can._status), can.isCmdMode() && can.onappend.style(can, html.OUTPUT)
		can.ui = can.onappend.layout(can, [[aaa.USERNAME, "mailbox", html.PROJECT], [html.PROFILE, html.CONTENT]]), can.onmotion.toggle(can, can.ui.profile, true)
		can.require([can.user.template(can, "trans.js")], function() { can.onimport._username(can, msg, can.ui.username) })
	},
	_table: function(can, msg, target, cb) { var _select
		msg.Table(function(value) { var action = can.page.parseAction(can, value).concat("", can.base.Obj(msg.Option(ice.MSG_ACTION)))
			_select = cb(value, function(event, cb) { can.user.carteRight(event, can, {}, action, cb) }, _select)
		}), can.onmotion.select(can, target, "", _select, function(target) { target.click() })
	},
	_username: function(can, msg, target) { can.onimport._table(can, msg, target, function(value, cb, _select) {
		value.nick = value.username, value.icon = can.onimport._user_icon(can, value.username)
		var _target = can.onimport.item(can, value, function(event) {
			if (can.onmotion.cache(can, function() { return value.username }, can.ui.mailbox)) {
				return can.page.Select(can, can.ui.mailbox, html.DIV_ITEM_SELECT, function(target) { target.click() })
			} can.run(event, [value.username], function(msg) { can.onimport._mailbox(can, msg, can.ui.mailbox, value.username) })
		}, function(event) { cb(event, function(event, button) { can.Update(can.request(event, value), [ctx.ACTION, button, value.username]) }) }, target)
		can.page.Append(can, _target, [{icon: "bi bi-arrow-repeat", style: {float: "right"}, onclick: function(event) {
			can.runAction(event, "sync", [value.username])
		}}])
		return can.db.hash[0] == value.username? _target: _select||_target
	}) },
	_mailbox: function(can, msg, target, username) { can.onimport._table(can, msg, target, function(value, cb, _select) {
		if (can.base.isIn(value.mailbox, "Drafts", "Spam", "Trash") && !value.count) { return _select }
		value.nick = value.username, value.icon = can.onimport._box_icon(can, value.mailbox)
		value.nick = value.order > 99? value.mailbox: can.user.trans(can, value.mailbox, can.onimport._box_trans(can))
		value.nick += lex.SP+value.count
		var _target = can.onimport.item(can, value, function(event) { can.page.ClassList.del(can, can.ui.content, html.SELECT)
			if (can.onmotion.cache(can, function() { return [username, value.mailbox].join(nfs.PS) }, can.ui.project)) {
				return can.page.Select(can, can.ui.project, html.DIV_ITEM_SELECT, function(target) { target.click() })
			} can.run(event, [username, value.mailbox], function(msg) { can.onimport._project(can, msg, can.ui.project, username, value.mailbox) })
		}, function(event) { cb(event, function(event, button) { can.Update(can.request(event, value, {username: username}), [ctx.ACTION, button, username, value.mailbox]) }) }, target)
		return can.db.hash[0] == username && can.db.hash[1] == value.mailbox? _target: _select||_target
	}) },
	_project: function(can, msg, target, username, mailbox) {
		if (msg.Length() == 0) { return can.onmotion.cache(can, function() { return [username, mailbox, "0"].join(nfs.PS) }, can.ui.profile) }
		can.onimport.filter(can, target).placeholder = `search in ${msg.Length()} items`
		can.onimport._table(can, msg, target, function(value, cb, _select) {
			var _target = can.page.Append(can, target, [{view: html.ITEM, list: [
				{view: [nfs.SUBJECT, "", value.subject]}, {view: log.INFO, list: [
					{view: nfs.FROM, list: [{icon: can.onimport._user_icon(can, value.from)}, {text: "  "}, {text: can.page.trans(can, value.from)}]}, {view: web.SPACE},
					{view: [mdb.TIME, "", can.onexport.time(can, value.time)]}
				]},
			], onclick: function(event) { can.onmotion.select(can, target, "", event.currentTarget), can.misc.SearchHash(can, username, mailbox, value.id)
				if (can.user.isMobile && !can.user.isLandscape()) { can.onmotion.hidden(can, can.ui.project.parentNode) }
				if (can.onmotion.cache(can, function() { return [username, mailbox, value.id].join(nfs.PS) }, can.ui.profile)) {
					can.onmotion.select(can, can.ui.profile.parentNode, html.DIV_CONTENT, value._target)
					return can.ui.content = value._target, can.onimport.layout(can)
				} can.run(event, [username, mailbox, value.id], function(msg) { can.onimport._content(can, msg, can.ui.content, username, mailbox, value) })
			}, oncontextmenu: function(event) { cb(event, function(event, button) { can.Update(can.request(event, value, {mailbox: mailbox, username: username}), [ctx.ACTION, button]) }) }}])._target
			return can.db.hash[0] == username && can.db.hash[1] == mailbox && can.db.hash[2] == value.id? _target: _select||_target
		})
	},
	_content: function(can, msg, target, username, mailbox, value) {
		can.page.Append(can, can.ui.profile, [{view: [nfs.SUBJECT, "", value.subject],
			onclick: function() { can.user.isMobile && can.onmotion.toggle(can, can.ui.project.parentNode, true) },
			ondblclick: function() { can.page.ClassList.set(can, can.ui.profile, "small", !can.onmotion.toggle(can, can.ui.project.parentNode)), can.onimport.layout(can) },
		}, {view: log.INFO, list: can.core.List([nfs.FROM, nfs.TO, "cc"], function(item) {
			return can.onexport.info(can, {name: item, icon: can.onimport._user_icon(can, value[item]), text: value[item]})
		}).concat([can.onexport.info(can, {name: mdb.TIME, icon: "bi bi-clock", text: can.onexport.time(can, value.time)})])}])
		var content = can.page.Append(can, can.ui.profile.parentNode, [html.CONTENT])._target; content.innerHTML = msg.Result()
		can.onmotion.select(can, can.ui.profile.parentNode, html.DIV_CONTENT, content)
		can.ui.content = value._target = content, can.onimport.layout(can)
	},
	layout: function(can) { can.user.isMobile && can.page.style(can, document.body, html.OVERFLOW, html.HIDDEN)
		if (can.Option(mdb.ID)) { return can.page.Select(can, can._output, "", function(target) { can.page.style(can, target, html.HEIGHT, can.ConfHeight()-1, html.WIDTH, can.ConfWidth()) }) }
		can.page.styleHeight(can, can.ui.project.parentNode, can.ConfHeight()), can.page.styleHeight(can, can.ui.profile.parentNode, can.ConfHeight())
		var height = can.ui.profile.offsetHeight, width = can.ui.project.parentNode.offsetWidth, border = 1
		if (can.user.isMobile && !can.user.isLandscape()) { width = 0, can.page.ClassList.del(can, can._output, html.FLEX) }
		can.page.styleHeight(can, can.ui.project, can.ConfHeight() > 800? can.ConfHeight() - can.ui.username.offsetHeight - can.ui.mailbox.offsetHeight: "")
		can.page.style(can, can.ui.profile.parentNode, html.WIDTH, can.ConfWidth()-width-border)
		can.page.Select(can, can.ui.content, "", function(target) { can.page.style(can, target, html.HEIGHT, can.ConfHeight()-height-border, html.WIDTH, can.ConfWidth()-width-border) })
	},
}, [""])
Volcanos(chat.ONACTION, {_trans: {from: "来自", to: "收件", cc: "抄送", time: "时间"}})
Volcanos(chat.ONEXPORT, {
	info: function(can, item) { return item.text && {view: item.name, list: [{text: can.user.trans(can, item.name)+": "}, {icon: item.icon}, {text: " "}, {text: can.page.trans(can, item.text)}]} },
	// time: function(can, time) { return can.base.trimPrefix(can.base.trimSuffix(time, ".000"), new Date().getFullYear()+"-") },
	time: function(can, time) { return can.base.TimeTrim(time) },
})

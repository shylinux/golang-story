Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) {
		// can.onmotion.hidden(can, can._status), can.isCmdMode() && can.onappend.style(can, html.OUTPUT)
		can.ui = can.onappend.layout(can), can.onmotion.toggle(can, can.ui.display, true)
		can.base.Copy(can.ui, can.page.Appends(can, can.ui.project, ["username", "mailbox", "maillist"]), "username", "mailbox", "maillist")
		can.require([can.user.template(can, "trans.js")], function() { can.onimport._username(can, msg, can.ui.username) })
	},
	_table: function(can, msg, target, cb) { var _select
		msg.Table(function(value) { var action = can.page.parseAction(can, value).concat("", can.base.Obj(msg.Option(ice.MSG_ACTION)))
			_select = cb(value, function(event, cb) { can.user.carteRight(event, can, {}, action, cb) }, _select)
		}), can.onmotion.select(can, target, "", _select, function(target) { target.click() })
	},
	_username: function(can, msg, target) { can.onimport._table(can, msg, target, function(value, cb, _select) {
		value.nick = value.username, value.icon = can.onimport._user_icon(can, value.username)
		value.nick = [{text: [value.username, "", "username"]}, {icon: "bi bi-arrow-repeat", onclick: function(event) { can.runAction(event, "sync", [value.username]) }}]
		var _target = can.onimport.item(can, value, function(event) {
			if (can.onmotion.cache(can, function() { return value.username }, can.ui.mailbox)) {
				return can.page.Select(can, can.ui.mailbox, html.DIV_ITEM_SELECT, function(target) { target.click() })
			}
			can.run(event, [value.username], function(msg) {
				can.onimport._mailbox(can, msg, can.ui.mailbox, value.username)
			})
		}, function(event) {
			cb(event, function(event, button) { can.Update(can.request(event, value), [ctx.ACTION, button, value.username]) })
		}, target)
		return can.db.hash[0] == value.username? _target: _select||_target
	}) },
	_mailbox: function(can, msg, target, username) { can.onimport._table(can, msg, target, function(value, cb, _select) {
		if (can.base.isIn(value.mailbox, "Drafts", "Spam", "Trash") && !value.count) { return _select }
		value.nick = value.username, value.icon = can.onimport._box_icon(can, value.mailbox)
		value.nick = value.order > 99? value.mailbox: can.user.trans(can, value.mailbox, can.onimport._box_trans(can))
		value.nick += lex.SP+value.count
		var _target = can.onimport.item(can, value, function(event) { can.page.ClassList.del(can, can.ui.content, html.SELECT)
			if (can.onmotion.cache(can, function() { return [username, value.mailbox].join(nfs.PS) }, can.ui.maillist)) {
				return can.page.Select(can, can.ui.project, html.DIV_ITEM_SELECT, function(target) { target.click() })
			}
			can.run(event, [username, value.mailbox], function(msg) {
				can.onimport._maillist(can, msg, can.ui.maillist, username, value.mailbox)
			})
		}, function(event) {
			cb(event, function(event, button) { can.Update(can.request(event, value, {username: username}), [ctx.ACTION, button, username, value.mailbox]) })
		}, target)
		return can.db.hash[0] == username && can.db.hash[1] == value.mailbox? _target: _select||_target
	}) },
	_maillist: function(can, msg, target, username, mailbox) {
		if (msg.Length() == 0) { return can.onmotion.cache(can, function() { return [username, mailbox, "0"].join(nfs.PS) }, can.ui.profile) }
		can.onappend.filter(can, target).placeholder = `search in ${msg.Length()} items`
		can.onimport._table(can, msg, target, function(value, cb, _select) {
			var _target = can.page.Append(can, target, [{view: html.ITEM, list: [
				{view: [nfs.SUBJECT, "", value.subject]}, {view: log.INFO, list: [
					{view: nfs.FROM, list: [{icon: can.onimport._user_icon(can, value.from)}, {text: "  "}, {text: can.page.trans(can, value.from)}]}, {view: web.SPACE},
					{view: [mdb.TIME, "", can.onexport.time(can, value.time)]}
				]},
			], onclick: function(event) { can.onmotion.select(can, target, "", event.currentTarget), can.misc.SearchHash(can, username, mailbox, value.id)
				if (can.user.isMobile && !can.user.isLandscape()) { can.onmotion.hidden(can, can.ui.project.parentNode) }
				if (can.onmotion.cache(can, function() { return [username, mailbox, value.id].join(nfs.PS) }, can.ui.content)) {
					can.onmotion.select(can, can.ui.display.parentNode, html.DIV_DISPLAY, value._display)
					return can.ui.display = value._display, can.onimport.layout(can)
				}
				can.run(event, [username, mailbox, value.id], function(msg) {
					can.onimport._content(can, msg, can.ui.content, username, mailbox, value)
				})
			}, oncontextmenu: function(event) { cb(event, function(event, button) { can.Update(can.request(event, value, {mailbox: mailbox, username: username}), [ctx.ACTION, button]) }) }}])._target
			return can.db.hash[0] == username && can.db.hash[1] == mailbox && can.db.hash[2] == value.id? _target: _select||_target
		})
	},
	_content: function(can, msg, target, username, mailbox, value) {
		can.page.Append(can, target, [{view: [nfs.SUBJECT, "", value.subject],
			onclick: function() { can.user.isMobile && can.onmotion.toggle(can, can.ui.project.parentNode, true) },
			ondblclick: function() { can.page.ClassList.set(can, can.ui.content, "small", !can.onmotion.toggle(can, can.ui.project)), can.onimport.layout(can) },
		}, {view: log.INFO, list: [can.onexport.info(can, {name: mdb.TIME, icon: "bi bi-clock", text: can.onexport.time(can, value.time)})].concat(
			can.core.List([nfs.FROM, nfs.TO, "cc"], function(item) { return can.onexport.info(can, {name: item, icon: can.onimport._user_icon(can, value[item]), text: value[item]}) }),
		)}])
		var display = can.page.Append(can, can.ui.display.parentNode, [html.DISPLAY])._target; display.innerHTML = msg.Result()
		can.onmotion.select(can, can.ui.display.parentNode, html.DIV_DISPLAY, display)
		can.ui.display = value._display = display, can.onimport.layout(can)
	},
	layout: function(can) {
		can.page.style(can, can.ui.project, html.HEIGHT, can.ConfHeight())
		can.page.style(can, can.ui.display, html.HEIGHT, can.ConfHeight()-can.ui.content.offsetHeight)
	},
}, [""])
Volcanos(chat.ONACTION, {_trans: {from: "来自", to: "收件", cc: "抄送", time: "时间"}})
Volcanos(chat.ONEXPORT, {
	time: function(can, time) { return can.base.TimeTrim(time) },
	info: function(can, item) { return item.text && {view: [[html.ITEM, item.name]], list: [{text: can.user.trans(can, item.name)+": "}, {icon: item.icon}, {text: " "}, {text: can.page.trans(can, item.text)}]} },
})

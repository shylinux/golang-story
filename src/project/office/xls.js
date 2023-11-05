Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) { can.page.requireModules(can, ["xlsx/dist/xlsx.full.min.js"], async function() {
		const book = XLSX.read(await (await fetch(msg.Option(nfs.PATH))).arrayBuffer()); can._output._cache_key = ""
		can.onimport.tabs(can, book.SheetNames, function(event, tabs) { can.db.sheet = book.Sheets[tabs.name]
			if (can.onmotion.cache(can, function() { return tabs.name })) { return } can._output.innerHTML = XLSX.utils.sheet_to_html(can.db.sheet)
			can.page.Select(can, can._output, html.TABLE, function(target) { can.onappend.style(can, [html.CONTENT, "xls"], target) })
		})
	}) },
}, [""])
Volcanos(chat.ONACTION, {
	parse: function(event, can, button) { can.user.input(event, can, [nfs.SCRIPT], function(list) {
		debugger
		can.onappend.plugin(can, {title: can.user.trans(can, "parse", "解析"), display: web.SHARE_LOCAL+can._msg.Option(nfs.DIR_ROOT)+list[0], style: html.FLOAT, height: can.ConfHeight()*3/4, width: can.ConfWidth()*3/4, top: can.ConfHeight()/4, left: can.ConfWidth()/4}, function(sub) {
			sub.run = function(event, cmds, cb) {
				var msg = can.request(); msg._sheet = can.db.sheet, cb(msg)
				can.onmotion.resize(can, sub._target, function(height, width) { sub.onimport.size(sub, height, width, true) })
			}
		})
	}) },
})

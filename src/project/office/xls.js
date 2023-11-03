Volcanos(chat.ONIMPORT, {
	_init: function(can, msg) { can.page.requireModules(can, ["xlsx/dist/xlsx.full.min.js"], async function() { can._output._cache_key = ""
		const book = XLSX.read(await (await fetch(msg.Option(nfs.PATH))).arrayBuffer())
		can.onimport.tabs(can, book.SheetNames, function(event, tabs) { if (can.onmotion.cache(can, function() { return tabs.name })) { return }
			var sheet = book.Sheets[tabs.name]; can._output.innerHTML = XLSX.utils.sheet_to_html(sheet)
			can.page.Select(can, can._output, html.TABLE, function(target) { can.onappend.style(can, html.CONTENT, target) })
		})
	}) },
})

Volcanos(chat.ONIMPORT, {
	_box_trans: function(can) {
		return {
			"INBOX": "收件箱",
			"Sent": "已发送",
			"Drafts": "草稿箱",
			"Spam": "垃圾邮件",
			"Trash": "已删除",
			"Archive": "已归档",
		}
	},
	_box_icon: function(can, mail) {
		return {
			"INBOX": "bi bi-mailbox",
			"Sent": "bi bi-send",
			"Drafts": "bi bi-send-exclamation",
			"Spam": "bi bi-envelope-slash",
			"Trash": "bi bi-trash",
			"Archive": "bi bi-archive",
		}[mail]||"bi bi-folder"
	},
	_user_icon: function(can, user) {
		return {
			"Mail Delivery System<MAILER-DAEMON@mail.shylinux.com>": "bi bi-exclamation-circle",
			"Rocket.Chat Cloud<cloud@rocket.chat>": "bi bi-chat",
			"GitHub<noreply@github.com>": "bi bi-github",
		}[user]||"bi bi-person"
	},
}, ["/require/modules/bootstrap-icons/font/bootstrap-icons.min.css"])

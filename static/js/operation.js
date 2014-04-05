var ConfirmDelete = function (url, msg) {
	question = window.confirm(msg);
	if (question == true) {
		//alert(url);
		//window.location.href = url;
		//window.location.href = "/admin/news/delete/nid/11";
		//window.location.assign(url);
		window.open(url);
	}
}
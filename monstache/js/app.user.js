module.exports = function (doc) {
	delete doc.password;
	return doc;
};

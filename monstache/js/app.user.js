module.exports = function (doc) {
	doc._meta_monstache = {
		version: doc.version || new Date().getTime(),
	};

	delete doc.password;
	delete doc.version;
	return doc;
};

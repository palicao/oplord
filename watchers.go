package oplord

var Watchers = []Watcher{
	{
		Matcher: SimpleMatcherFactory("test.mycollection", []string{"u", "i"}),
		Action:  SimplePostActionFactory("http://localhost:8080/hook"),
	},
	{
		Matcher: SimpleMatcherFactory("test.mycollection", []string{"d"}),
		Action:  SimplePostActionFactory("http://localhost:8080/delete-hook"),
	},
}

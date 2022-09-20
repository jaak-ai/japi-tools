package road

type Road struct {
	Children []Route
}

func (r *Road) Configure(app RR) {
	r.buildRoute(app, r.Children)
}

func (r *Road) buildRoute(app RR, children []Route) {
	for _, route := range children {

		if len(route.Children) > 0 {
			group := app.Group(route.Path, route.Handler...)
			r.buildRoute(group, route.Children)
			continue
		}

		app.Add(route.Method, route.Path, route.Handler...)
	}
}

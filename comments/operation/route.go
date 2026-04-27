package operation

// Route defines a path and a method
type Route struct {
	Path   string
	Method string
}

// Routes is a type alias for a slice of Route
type Routes []Route

// Summarize returns a map of paths to methods for the operation.
func (r Routes) Summarize() map[string][]string {
	summary := make(map[string][]string)
	for _, route := range r {
		summary[route.Path] = append(summary[route.Path], route.Method)
	}
	return summary
}

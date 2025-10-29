package widget

// Widget is a widget that can be rendered. It's a simple interface that
// requires a single method, View().
type Widget interface {
	View() string
}

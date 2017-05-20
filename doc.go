/*
Package ataman is a colored terminal text rendering engine using ANSI sequences.
The project aims on simple text attribute manipulations with templates.

Here is the couple of examples to introduce the project.

    // Rendering with the renderer
    rndr := ataman.NewRenderer(ataman.CurlyStyle())
    tpl := "{red}Red {green}Green {blue}Blue{-}"
    fmt.Println(rndr.MustRender(tpl))

    // Rendering pre-rendered template.
    prep := rndr.MustPrepare("{light_green}%s{-}, {bg_light_yellow+blue+bold}%s{-}!")
    fmt.Println(prep.Format("Hello", "World"))

*/
package ataman


# lterm

lterm is a basic terminal color package for go.

## Example

```
bg := lterm.Bg8(240)
fg1 := lterm.Fg8(235)
fg2 := lterm.Fg8(228)

lterm.Printline(bg, " test", fg1, " foo ", fg2, "bar ", lterm.Reset)
```

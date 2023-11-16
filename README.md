# Usage

```
$ echo "Left :-: Middle :-: Right" | cutnstitch ' :-: ' '{{index . 0}} --> {{index . 2}}'
Left --> Right
```

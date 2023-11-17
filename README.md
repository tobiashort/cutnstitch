# Usage

```
$ cat <<'EOF' | cutnstitch ' :-: ' '{{index . 0}} --> {{index . 2}}'
Left :-: Middle :-: Right
Left :-: Middle :-: Right
Left :-: Middle :-: Right
EOF
Left --> Right
Left --> Right
Left --> Right
```

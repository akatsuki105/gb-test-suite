# ei_sequence

連続でEIを実行した時の挙動についてチェックします。

```asm
.include "common.s"

  di
  ld a, INTR_SERIAL
  ld (IF), a
  ld (IE), a
  xor a
  ld b, a
  ld c, a

  jr test

.org $1A0
test:
  ei
.repeat 17
  ei
.endr

fail:
  quit_failure_string "FAIL: NO INTR"

test_finish:
  setup_assertions
  assert_b $01
  assert_c $A2
  quit_check_asserts

.org INTR_VEC_SERIAL
  pop bc
  jp test_finish
```
# halt_ime0_nointr_timing

このテストでは、`IME=0`でHALTをしたときに追加で遅延が起きるかどうかをテストします。

`IME=0`の場合、HALTは、割り込みを待つために長い一連のNOP命令を使用した場合と全く同じタイミングで、直ちに実行を継続することが期待されます。

```asm
.include "common.s"

.macro clear_IF
  xor a
  ldh (<IF), a
.endm

.macro enable_IE_vblank
  ld a, INTR_VBLANK
  ldh (<IE), a
.endm

  di
  wait_ly 10
  enable_IE_vblank

; ----------- ROUND 1 -----------

  clear_IF
  ld hl, test_round1

  ; VBlankで test_round1 にジャンプするのを待つ
  ei
  halt  ; このとき、ie & if & 0x1f == 0 なのでHALTします
  nop
  jp fail_halt

test_round1:
  ld hl, fail_intr
  clear_IF

  nops 13
  xor a
  ldh (<DIV), a

  halt      ; IME=0 なのでHALTしない (VBlankでRETIしていなので IME=0 に注意)
  nops 6    ; Equivalent to interrupt + JP HL in the IME=1 case

finish_round1:
  ldh a, (<DIV)
  ld d, a

; ----------- ROUND 2 -----------

  clear_IF
  ld hl, test_round2

  ei
  halt
  nop
  jp fail_halt

test_round2:
  ld hl, fail_intr
  clear_IF

  nops 12
  xor a
  ldh (<DIV), a

  halt
  nops 6 ; Equivalent to interrupt + JP HL in the IME=1 case

finish_round2:
  ldh a, (<DIV)
  ld e, a

; ----------- Result -----------

  setup_assertions
  assert_d $11  ; ROUND1のDIVの値
  assert_e $12  ; ROUND2のDIVの値
  quit_check_asserts

fail_halt:
  quit_failure_string "FAIL: HALT"

fail_intr:
  quit_failure_string "FAIL: INTERRUPT"

; VBlankでHLの示すアドレスにジャンプ
; RETIはしないのでIME=0のまま
.org INTR_VEC_VBLANK
  jp hl
```
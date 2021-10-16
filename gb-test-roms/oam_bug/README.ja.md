# OAMメモリ破壊バグテスト

- DMGのOAMメモリ破壊バグを検証します。

- Occurs when 16-bit increment/decrement is made of value in range $FE00
to $FEFF, during around the first 20 cycles of a visible scanline while
LCD is on, where 114 cycles = 1 scanline.

- 数バイトのOAMをある場所から別の場所にコピーさせます。

- Occurs with instructions that do increment:

```asm
	INC rp (including SP)
	DEC rp
	POP rp      counts as two increments
	PUSH rp     counts as two increments
	LD A,(HL+)
	LD A,(HL-)
```

- Doesn't occur with instructions that do 16-bit add:

```asm
	LD HL,SP+n
	ADD HL,rp
	ADD SP,n
```

- 10回あるVBlankのスキャンラインの間、いつでも発生するわけではありません。

- LCD無効時には発生せず、有効時にはいつ発生してもおかしくありません。

- メモリ破壊でどのようにメモリが破壊されるかは、発生するタイミングに依存します。


## Multi-ROM

このディレクトリには、すべてのテストを一括で実行してくれる1つのテストROMがあります。(`oam_bug.gb`)

テストの番号を表示してテストを実行し、合格した場合は`ok`、そうでない場合はエラーコードを表示します。

すべてのテストが完了すると、すべてのテストが合格したことを報告するか、最初に失敗したテストの番号(1から数える)をエラーコードとして報告します。

テストが全部終わるとビープ音がなります。

テストが失敗しても、ディレクトリ`rom_singles`で対応するROM/GBSを見つけることで、テストを単独で実行することができます。

結果の画面がコンパクトになっているのは、結果が上にスクロールして前の結果が見えなくならないようにするためです。

## エラー情報

エラーコードや出力される情報の詳細については、`source/`にあるテストのソースコードを参照してください。

エラーコードNを見つけるには、`set_test N`でコードを検索してください。この`set_test N`は通常は失敗したサブテストの前に行われる処理です。

## Flashes, clicks, other glitches

一部のテストでは、画面にチラつきがあったり、わずかなノイズ音が発生する場合があります。

これはバグではありませんので、無視してください。特に断りのない限り、最後に表示されるテスト結果のみが重要です。

## LCD support

大抵の場合、テストは得られた情報を画面に出力するようになっています。

LCDをサポートしていないエミュレータや、固有のスクリーンを持たないGBSで実行しても、テストは問題なく行われます。
特に、VBlank待機ルーチンは、LYが現在のLCDラインを反映していないようなエミュレータに対するタイムアウト処理を持っています。

また、液晶ディスプレイがスクロールに対応していない場合でも、結果はちゃんと表示可能です。

## メモリへのテスト結果の書き込みについて

[mem_timing-2](../mem_timing-2/README.ja.md#メモリへのテスト結果の書き込みについて)参照

## GBS versions

Many GBS-based tests require that the GBS player either not interrupt
the init routine with the play routine, or if they do, not interrupt the
play routine again if it hasn't returned yet. This is because many tests
need to run for a while without returning.

In addition to the other text output methods described above, GBS builds
report essential information bytes audibly, including the final result.
A byte is reported as a series of tones. The code is in binary, with a
low tone for 0 and a high tone for 1. The first tone is always a zero. A
final code of 0 means passed, 1 means failure, and 2 or higher indicates
a specific reason as listed in the source code by the corresponding
set_code line. Examples:

Tones        | Binary | Decimal | Meaning
-----------  | ------ | ------- | -------
low          |  0     | 0       | passed
low high     |  01    | 1       | failed
low high low | 010    | 2       | error 2

-- 
Shay Green <gblargg@gmail.com>

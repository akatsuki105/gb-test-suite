# GameBoyのCPUによるメモリアクセス時間のテスト

このROMは、スタックやプログラムカウンタへのアクセスを除く、CPU命令によるメモリの読み書きのタイミングをテストします。

これらのテストでは、正しい命令のタイミングと適切なタイマー操作（TAC、TIMA、TMA）が要求されます。

読み書きテストでは、失敗した命令を以下のようにリストアップしています。

```
[CB] opcode:tested-correct
```

read-modify-writeテストでは、失敗した命令を以下のようにリストアップします。

```
[CB] opcode:tested read/tested write-correct read/correct write
```

オペコードの後の値は、アクセスがどの命令サイクルで行われるかを示しており、1が最初の命令サイクルとなります。

他の問題でアクセスのサイクルが確定できなかった場合は0と表示されます。

読み出しか書き込みのどちらかで、両方ではない命令の場合、CPUは最後のサイクルでアクセスを行います。

読み取り、修正、書き戻しを行う命令では、CPUは最後の1個前のサイクルで読み取り、最後のサイクルで書き込みを行います。

## テスト内部の処理について

[mem_timing](../mem_timing/README.ja.md#テスト内部の処理について)参照

## Multi-ROM

このディレクトリには、すべてのテストを一括で実行してくれる1つのテストROMがあります。(`mem_timing.gb`)

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

## Output to memory

テキスト出力と最終結果も`$A000`のメモリに書き込まれるので、CPUとRAMしかサポートしていない非常にミニマルなエミュレータでもテストを実行することができます。

ランダムなデータではなく、テストからのデータであることを確実に示すために、`$A001..$A003`にはマジックナンバ`$DE,$B0,$61`が書き込まれます。

このマジックナンバがあると、テキスト文字列と最終結果のステータスが有効になります。

`$A000` は、全体のステータスを保持しています。テストがまだ実行されている場合は、`0x80`を、そうでない場合は、最終結果コードを保持しています。

すべてのテキスト出力は、`$A004`のゼロ終端文字列に付加されます。

All text output is appended to a zero-terminated string at $A004. 

エミュレータでは、この文字列に追加の文字がないか定期的にチェックして、それを出力することで、最終的な結果を最後に出力するだけでなく、リアルタイムにテストで得られた情報を出力することができます。

## GBS versions

> GBS: Game Boy Sound Systemの略。GBSは、ゲームボーイサウンドハードウェア用に設計された音楽を再生するためのファイル形式です。

Many GBS-based tests require that the GBS player either not interrupt the init routine with the play routine, or if they do, not interrupt the play routine again if it hasn't returned yet. This is because many tests need to run for a while without returning.

In addition to the other text output methods described above, GBS builds report essential information bytes audibly, including the final result.
A byte is reported as a series of tones. The code is in binary, with a low tone for 0 and a high tone for 1. The first tone is always a zero.
A final code of 0 means passed, 1 means failure, and 2 or higher indicates a specific reason as listed in the source code by the corresponding set_code line.
Examples:

Tones        | Binary | Decimal | Meaning
-----------  | ------ | ------- | -------
low          |  0     | 0       | passed
low high     |  01    | 1       | failed
low high low | 010    | 2       | error 2


# <p align="center">TypeLang<p>

<p align="center">
<img src="./docs/assets/ton_green.svg" width="200">
<img src="./docs/assets/filler.svg" width="100">
<img src="./docs/assets/telegram_green.svg" width="200">
</p>

<!-- markdownlint-disable MD013 -->
<!-- it's unable to past links as references, when you trying to center links-->
<p align="center">
<a href="https://pkg.go.dev/github.com/xelaj/tl">
<img src="https://img.shields.io/badge/go.dev-reference-007d9c?label=%e2%a0%80&logo=go&logoColor=white&labelColor=27303B&style=for-the-badge">
</a>
<a href="https://goreportcard.com/report/github.com/xelaj/tl">
<img src="https://img.shields.io/static/v1?label=go+report&message=A%2b&color=success&labelColor=27303B&style=for-the-badge">
</a>
<a href="https://codecov.io/gh/xelaj/tl">
<img src="https://img.shields.io/codecov/c/gh/xelaj/mtproto?labelColor=27303B&label=cover&logo=codecov&style=for-the-badge">
</a>
<a href="https://bit.ly/2xlsVsQ">
<img src="https://img.shields.io/badge/chat-telegram-0088cc?labelColor=27303B&logo=telegram&style=for-the-badge">
</a>
<a href="https://github.com/xelaj/tl/blob/main/LICENSE.md">
<img src="https://img.shields.io/badge/license-MIT-green?labelColor=27303B&style=for-the-badge">
</a>
<br/>
<a href="https://github.com/xelaj/tl/releases">
<img src="https://img.shields.io/github/v/tag/xelaj/tl?labelColor=27303B&label=version&sort=semver&style=for-the-badge">
</a>
<img src="https://img.shields.io/static/v1?label=stability&message=stable&labelColor=27303B&color=success&style=for-the-badge">
<img src="https://img.shields.io/static/v1?label=%d1%81%d0%bb%d0%b0%d0%b2%d0%b0&message=%d0%a3%d0%ba%d1%80%d0%b0%d1%97%d0%bd%d1%96&color=ffd700&labelColor=0057b7&style=for-the-badge">
<!--
code quality
golangci
contributors
go version
gitlab pipelines
-->
</p>
<!-- markdownlint-enable MD013 -->

<p align="center">
Full native implementation of TL serializing language. Simple and fast.
</p>

**english** [русский][inex_ru]

## ✨ Features

TODO

## 👨‍💻 How to use

TODO

**Code examples are [here][gh_examples]**

**Full docs are [here][godoc]**

## ⌚️ Getting started

TODO

### Why TL serializing is **SO** ugly?

Well... Read [this issue][ton_issue] about TON source code, when it was released
first time. Remember: current TON developers are absolutely another guys. Use
google translate, this issue will answer to all your questions.

Telegram devs are really strange

## 🤔 Differences between TL and TL-B

Вы наверняка заметили, что логотип этого проекта это на самом деле два логотипа — telegram и ton. Однако, формат в ton обратно несовместим с тем же форматом в telegram. При этом, между ними очень много общего, поэтому вцелом довольно реально объеденить оба формата в одну библиотеку. При этом, в отличии от разработчиков telegram, разработчики ton более благосклонны к кооперации, так что вполне возможно, что поддержка ton будет даже более качественной

https://github.com/ton-blockchain/wallet-android/blob/-/app/jni/ton/crypto/block/block-auto.cpp

генератор

https://github.com/ton-blockchain/ton/blob/-/crypto/tl/tlbc-gen-cpp.cpp

## 🦊 Who use it

Mmmm...me?

## 💎 Contributing

Please read [contributing guide][gh_contributing] if you want to help. And the
help is very necessary!

**Don't want code?** Read [this page][gh_support]! We love nocoders!

## 🐛 Security bugs?

Please, **don't** create issue which describes security bug, this can be too
offensive! Instead, please read [this notification][gh_security] and follow that
steps to notify us about problem.

## 🏋️ TODO

- [x] Full implementation of Tl serializing
- [x] Tool to generate golang code
- [ ] Write amazing docs
- [ ] Create full toolset to support TL
- [ ] make support of TL-B (incompatible version, used by TON)
- [ ] Write RFC specification?
- [ ] write good security policy
- [ ] setup taskfile for markdown linter

## 📒 Running project scripts

This project uses [go-task][taskfile], it's not important to undestand, what's
going on, (since you can just see into `Taskfile.yaml` and see all commands).
For better experience, you can download go-task and run tasks e.g. via
`$ task <taskname>`. All tasks can be shown via `$ task --list-all`

## 👨‍👩‍👧‍👦 Authors

- **Richard Cooper** <[rcooper.xelaj@protonmail.com](mailto:rcooper.xelaj@protonmail.com)>

## 📝 License

This project is licensed under the MIT License - see the [LICENSE][license_en]
file for details

Если вы находитесь в россии, или как-либо связаны с российским правительством,
(например, являетесь российским налогоплательщиком) на вас распостраняется
[отдельная лицензия][license_ru].

## One important thing

Even that maintainers of this project are generally from russia, we still stand
up with Ukraine, and from beginning of war, decided to stop paying any taxes, or
cooperate in any case with government, and companies, connected with government.
This is absolutely nothing compared to how much pain putin brought to the
fraternal country. And we are responsible for our inaction, and the only thing
we can do is to take at least any actions that harm putin’s regime, and help the
victims of regime using all resources available for us.
<img src="./docs/assets/by_flag.svg" height="16">
<img src="./docs/assets/ru_flag.svg" height="16">
<img src="./docs/assets/ua_flag.svg" height="16">

<!--
V2UndmUga25vd24gZWFjaCBvdGhlciBmb3Igc28gbG9uZwpZb3
VyIGhlYXJ0J3MgYmVlbiBhY2hpbmcgYnV0IHlvdSdyZSB0b28g
c2h5IHRvIHNheSBpdApJbnNpZGUgd2UgYm90aCBrbm93IHdoYX
QncyBiZWVuIGdvaW5nIG9uCldlIGtub3cgdGhlIGdhbWUgYW5k
IHdlJ3JlIGdvbm5hIHBsYXkgaXQKQW5kIGlmIHlvdSBhc2sgbW
UgaG93IEknbSBmZWVsaW5nCkRvbid0IHRlbGwgbWUgeW91J3Jl
IHRvbyBibGluZCB0byBzZWU=
-->

--------------------------------------------------------------------------------

<p align=center><sub><sub>
Created with love 💜 and magic 🦄 </br> Xelaj Software, 2022
</sub></sub></p>

[taskfile]:      https://taskfile.dev/
[ton_issue]:     https://github.com/ton-blockchain/ton/issues/31
[telegram_chat]: https://t.me/xelaj_developers

<!-- images -->

[goreport_card]: https://goreportcard.com/badge/github.com/xelaj/tl

<!-- localizations -->
[inex_ru]: https://github.com/xelaj/tl/blob/-/docs/ru_RU/README.md

<!-- project links -->
[godoc]:              https://pkg.go.dev/github.com/xelaj/tl
[license_ru]:         https://github.com/xelaj/tl/blob/-/docs/ru_XZ/LICENSE.md
[license_en]:         https://github.com/xelaj/tl/blob/-/LICENSE.md
[gh_examples]:        https://github.com/xelaj/tl/blob/-/examples
[gh_security]:        https://github.com/xelaj/tl/blob/-/.github/SECURITY.md
[gh_support]:         https://github.com/xelaj/tl/blob/-/.github/SUPPORT.md
[gh_contributing]:    https://github.com/xelaj/tl/blob/-/.github/CONTRIBUTING.md
[gh_project]:         https://github.com/xelaj/tl/projects
[gh_discussions]:     https://github.com/xelaj/tl/discussions
[gh_discussions_faq]: https://github.com/xelaj/tl/discussions/categories/q-a
[CoC]:                https://github.com/xelaj/tl/blob/-/.github/CODE_OF_CONDUCT.md

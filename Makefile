default: run
_defaultOS = linux
_defaultARCH = amd64

program_name = discordRelativeTimestamp

targetOS = $(_defaultOS)
targetARCH = $(_defaultARCH)

_set_OS_ARCH:
	go env -w GOOS=$(targetOS)
	go env -w GOARCH=$(targetARCH)

run:
	go run .

buildALL: | buildp buildd

buildd: | _set_OS_ARCH
	go build -o ./EXPORT/$(program_name)_dev_$(targetOS)_$(targetARCH) ./main.go

buildp: | _set_OS_ARCH
	go build -ldflags "-w -s" -o ./EXPORT/$(program_name)_prod_$(targetOS)_$(targetARCH) ./main.go


# Для разных библиотек.
## Fyne
fyne-win:
	fyne package -os windows -icon ./static/icon.png -release --id Medtereus.discrodRelativeTimestamp

fyne-lin-p:
	fyne package -os linux -icon ./static/icon.png -release --id Medtereus.discrodRelativeTimestamp

fyne-lin-b:
	fyne build -os linux

# fyne-apk:
# 	fyne package -os android -app-id "Medtereus.$(program_name)" icon icon.png

# go env
set_linuxCC:
	go env -w CC=gcc
	go env -w CXX=g++

set_winCC:
	go env -w CC=x86_64-w64-mingw32-gcc
	go env -w CXX=x86_64-w64-mingw32-g++

gitC:
	git add .
	git commit -m "commit from Makefile"
	git push -u origin main
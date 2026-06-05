# Discord Relative Timestamp
[🇬🇧 English README](./README.md)
## Description
Discord supports a special timestamp format:

```text
<t:UNIX_TIMESTAMP:R>
```

After sending such a tag, Discord automatically displays relative time:

```text
in 2 days
in 5 hours
3 minutes ago
5 hours ago
```

I got tired of manually calculating Unix Timestamps using online converters every time, so I made my own application for this task.
Enter the number of days, hours, minutes, and seconds, press Copy — and the generated Discord tag will be copied to the clipboard.

| Main Screen                    | Settings                        |
| ------------------------------ | ------------------------------- |
| ![main](./doc/main_screen.png) | ![settings](./doc/settings.png) |

### Features
* generate Discord relative timestamps;
* generate timestamps for both future and past time;
* add time relative to the current moment;
* one-click copy to clipboard;
* light and dark themes;
* UI Scale setting;
* minimalist interface;
* works locally without an internet connection.

## Running
Download the binary from Releases and run it.
The application does not require installation and works locally.

## Technologies Used
* golang
* fyne

## Version History
* 0.0: Project created.
* 0.1: Double ESC press closes the application.
* 0.2: Everything is ready. Just a bit of polishing before release.
* 1.0: Release.

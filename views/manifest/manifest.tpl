<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>items</key>
        <array>
            <dict>
                <key>assets</key>
                <array>
                    <dict>
                        <key>kind</key>
                        <string>software-package</string>
                        <key>url</key>
                        <string>{{ .manifest.FileURL }}</string>
                    </dict>
                    <dict>
                        <key>kind</key>
                        <string>display-image</string>
                        <key>url</key>
                        <string>{{ .manifest.SmallImageURL }}</string>
                    </dict>
                    <dict>
                        <key>kind</key>
                        <string>full-size-image</string>
                        <key>url</key>
                        <string>{{ .manifest.FullImageURL }}</string>
                    </dict>
                </array>
                <key>metadata</key>
                <dict>
                    <key>bundle-identifier</key>
                    <string>{{ .manifest.Identifier }}</string>
                    <key>bundle-version</key>
                    <string>{{ .manifest.Version }}</string>
                    <key>kind</key>
                    <string>software</string>
                    <key>title</key>
                    <string>{{ .manifest.Title }}</string>
                </dict>
            </dict>
        </array>
    </dict>
</plist>

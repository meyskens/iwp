This needs dependencies:

sudo apt-get install libwebkit2gtk-4.0-dev
go build -tags gtk_3_10
sudo ln -s /usr/lib/x86_64-linux-gnu/pkgconfig/javascriptcoregtk-4.0.pc /usr/lib/x86_64-linux-gnu/pkgconfig/javascriptcoregtk-3.0.pc
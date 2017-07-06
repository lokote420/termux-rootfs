/*
The MIT License (MIT)

Copyright (c) 2017 Leonid Plyushch <leonid.plyushch@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package mime_types

import (
	"mime"
)

func initApplicationMime() {
	mime.AddExtensionType(".32x", "application/x-genesis-32x-rom")
	mime.AddExtensionType(".a", "application/x-archive")
	mime.AddExtensionType(".abw", "application/abiword")
	mime.AddExtensionType(".ace", "application/x-ace")
	mime.AddExtensionType(".agb", "application/x-gba-rom")
	mime.AddExtensionType(".ai", "application/illustrator")
	mime.AddExtensionType(".apk", "application/vnd.android.package-archive")
	mime.AddExtensionType(".ar", "application/x-archive")
	mime.AddExtensionType(".arj", "application/x-arj")
	mime.AddExtensionType(".asp", "application/x-asp")
	mime.AddExtensionType(".awk", "application/x-awk")
	mime.AddExtensionType(".axs", "application/olescript")
	mime.AddExtensionType(".bcpio", "application/x-bcpio")
	mime.AddExtensionType(".blend", "application/x-blender")
	mime.AddExtensionType(".bsdiff", "application/x-bsdiff")
	mime.AddExtensionType(".bz2", "application/x-bzip2")
	mime.AddExtensionType(".cab", "application/vnd.ms-cab-compressed")
	mime.AddExtensionType(".cap", "application/vnd.tcpdump.pcap")
	mime.AddExtensionType(".cat", "application/vnd.ms-pkiseccat")
	mime.AddExtensionType(".cct", "application/x-director")
	mime.AddExtensionType(".cdf", "application/cdf")
	mime.AddExtensionType(".cer", "application/x-x509-ca-cert")
	mime.AddExtensionType(".cfc", "application/x-cfm")
	mime.AddExtensionType(".cfm", "application/x-cfm")
	mime.AddExtensionType(".class", "application/x-java")
	mime.AddExtensionType(".clp", "application/x-msclip")
	mime.AddExtensionType(".com", "application/x-ms-dos-executable")
	mime.AddExtensionType(".core", "application/x-core")
	mime.AddExtensionType(".cpio", "application/x-cpio")
	mime.AddExtensionType(".crd", "application/x-mscardfile")
	mime.AddExtensionType(".crl", "application/pkix-crl")
	mime.AddExtensionType(".crt", "application/x-x509-ca-cert")
	mime.AddExtensionType(".csr", "application/pkcs10")
	mime.AddExtensionType(".cst", "application/x-director")
	mime.AddExtensionType(".cxt", "application/x-director")
	mime.AddExtensionType(".dar", "application/x-dar")
	mime.AddExtensionType(".dcr", "application/x-director")
	mime.AddExtensionType(".der", "application/x-x509-ca-cert")
	mime.AddExtensionType(".desktop", "application/x-desktop")
	mime.AddExtensionType(".dir", "application/x-director")
	mime.AddExtensionType(".dll", "application/x-sharedlib")
	mime.AddExtensionType(".dmg", "application/x-apple-diskimage")
	mime.AddExtensionType(".dmp", "application/vnd.tcpdump.pcap")
	mime.AddExtensionType(".doc", "application/vnd.ms-word")
	mime.AddExtensionType(".docbook", "application/x-docbook+xml")
	mime.AddExtensionType(".docm", "application/vnd.ms-word.document.macroEnabled.12")
	mime.AddExtensionType(".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	mime.AddExtensionType(".dot", "application/msword")
	mime.AddExtensionType(".dotm", "application/vnd.ms-word.template.macroEnabled.12")
	mime.AddExtensionType(".dotx", "application/vnd.openxmlformats-officedocument.wordprocessingml.template")
	mime.AddExtensionType(".dta", "application/x-stata")
	mime.AddExtensionType(".dtd", "application/xml-dtd")
	mime.AddExtensionType(".dvi", "application/x-dvi")
	mime.AddExtensionType(".dxf", "application/x-autocad")
	mime.AddExtensionType(".dxr", "application/x-director")
	mime.AddExtensionType(".elc", "application/x-elc")
	mime.AddExtensionType(".elf", "application/x-executable")
	mime.AddExtensionType(".enl", "application/x-endnote-library")
	mime.AddExtensionType(".enz", "application/x-endnote-library")
	mime.AddExtensionType(".eot", "application/vnd.ms-fontobject")
	mime.AddExtensionType(".eps", "application/postscript")
	mime.AddExtensionType(".epub", "application/epub+zip")
	mime.AddExtensionType(".etheme", "application/x-e-theme")
	mime.AddExtensionType(".evy", "application/envoy")
	mime.AddExtensionType(".exe", "application/x-ms-dos-executable")
	mime.AddExtensionType(".fb2", "application/x-fictionbook+xml")
	mime.AddExtensionType(".fif", "application/fractals")
	mime.AddExtensionType(".fl", "application/x-fluid")
	mime.AddExtensionType(".fm", "application/vnd.framemaker")
	mime.AddExtensionType(".fqd", "application/x-director")
	mime.AddExtensionType(".gba", "application/x-gba-rom")
	mime.AddExtensionType(".gem", "application/x-tar")
	mime.AddExtensionType(".glade", "application/x-glade")
	mime.AddExtensionType(".gnc", "application/x-gnucash")
	mime.AddExtensionType(".gnucash", "application/x-gnucash")
	mime.AddExtensionType(".gnumeric", "application/x-gnumeric")
	mime.AddExtensionType(".gnuplot", "application/x-gnuplot")
	mime.AddExtensionType(".gpg", "application/pgp-encrypted")
	mime.AddExtensionType(".gplt", "application/x-gnuplot")
	mime.AddExtensionType(".gpx", "application/gpx+xml")
	mime.AddExtensionType(".gra", "application/x-graphite")
	mime.AddExtensionType(".gsf", "application/x-font-type1")
	mime.AddExtensionType(".gtar", "application/tar")
	mime.AddExtensionType(".gz", "application/gzip")
	mime.AddExtensionType(".hdd", "application/x-virtualbox-hdd")
	mime.AddExtensionType(".hdf", "application/x-hdf")
	mime.AddExtensionType(".hlp", "application/winhlp")
	mime.AddExtensionType(".hqx", "application/binhex")
	mime.AddExtensionType(".hta", "application/hta")
	mime.AddExtensionType(".icm", "application/vnd.iccprofile")
	mime.AddExtensionType(".iii", "application/x-iphone")
	mime.AddExtensionType(".img", "application/x-raw-disk-image")
	mime.AddExtensionType(".indd", "application/x-indesign")
	mime.AddExtensionType(".ins", "application/x-internet-signup")
	mime.AddExtensionType(".iso", "application/x-cd-image")
	mime.AddExtensionType(".isp", "application/x-internet-signup")
	mime.AddExtensionType(".jar", "application/java-archive")
	mime.AddExtensionType(".jks", "application/x-java-keystore")
	mime.AddExtensionType(".jnlp", "application/x-java-jnlp-file")
	mime.AddExtensionType(".jsm", "application/javascript")
	mime.AddExtensionType(".json", "application/json")
	mime.AddExtensionType(".kml", "application/vnd.google-earth.kml+xml")
	mime.AddExtensionType(".kmz", "application/vnd.google-earth.kmz")
	mime.AddExtensionType(".ko", "application/x-object")
	mime.AddExtensionType(".ks", "application/x-java-keystore")
	mime.AddExtensionType(".la", "application/x-shared-library-la")
	mime.AddExtensionType(".latex", "application/x-latex")
	mime.AddExtensionType(".lha", "application/x-lha")
	mime.AddExtensionType(".lhz", "application/x-lhz")
	mime.AddExtensionType(".lib", "application/x-endnote-library")
	mime.AddExtensionType(".llb", "application/x-labview")
	mime.AddExtensionType(".lnk", "application/x-ms-shortcut")
	mime.AddExtensionType(".lrz", "application/x-lrzip")
	mime.AddExtensionType(".lvx", "application/x-labview-exec")
	mime.AddExtensionType(".lz4", "application/x-lz4")
	mime.AddExtensionType(".lz", "application/x-lzip")
	mime.AddExtensionType(".lzh", "application/x-lha")
	mime.AddExtensionType(".lzma", "application/x-lzma")
	mime.AddExtensionType(".lzo", "application/x-lzop")
	mime.AddExtensionType(".m4", "application/x-m4")
	mime.AddExtensionType(".ma", "application/mathematica")
	mime.AddExtensionType(".man", "application/x-troff-man")
	mime.AddExtensionType(".mbox", "application/mbox")
	mime.AddExtensionType(".mcd", "application/x-mathcad")
	mime.AddExtensionType(".mdb", "application/vnd.ms-access")
	mime.AddExtensionType(".me", "application/x-troff-me")
	mime.AddExtensionType(".mfp", "application/x-shockwave-flash")
	mime.AddExtensionType(".mht", "application/x-mimearchive")
	mime.AddExtensionType(".mhtml", "application/x-mimearchive")
	mime.AddExtensionType(".mif", "application/vnd.framemaker")
	mime.AddExtensionType(".mny", "application/x-msmoney")
	mime.AddExtensionType(".mobi", "application/x-mobipocket-ebook")
	mime.AddExtensionType(".mpp", "application/vnd.ms-project")
	mime.AddExtensionType(".ms", "application/x-troff-ms")
	mime.AddExtensionType(".mvb", "application/x-msmediaview")
	mime.AddExtensionType(".mws", "application/x-maple")
	mime.AddExtensionType(".nb", "application/mathematica")
	mime.AddExtensionType(".nes", "application/x-nes-rom")
	mime.AddExtensionType(".nez", "application/x-nes-rom")
	mime.AddExtensionType(".ngp", "application/x-neo-geo-pocket-rom")
	mime.AddExtensionType(".o", "application/x-object")
	mime.AddExtensionType(".obt", "application/x-openbox-theme")
	mime.AddExtensionType(".oda", "application/oda")
	mime.AddExtensionType(".odf", "application/vnd.oasis.opendocument.formula")
	mime.AddExtensionType(".odg", "application/vnd.oasis.opendocument.graphics")
	mime.AddExtensionType(".odp", "application/vnd.oasis.opendocument.presentation")
	mime.AddExtensionType(".ods", "application/vnd.oasis.opendocument.spreadsheet")
	mime.AddExtensionType(".odt", "application/vnd.oasis.opendocument.text")
	mime.AddExtensionType(".one", "application/msonenote")
	mime.AddExtensionType(".otf", "application/x-font-otf")
	mime.AddExtensionType(".ova", "application/x-virtualbox-ova")
	mime.AddExtensionType(".ovf", "application/x-virtualbox-ovf")
	mime.AddExtensionType(".p10", "application/pkcs10")
	mime.AddExtensionType(".p12", "application/x-x509-ca-cert")
	mime.AddExtensionType(".p7b", "application/pkcs7-mime")
	mime.AddExtensionType(".p7c", "application/pkcs7-mime")
	mime.AddExtensionType(".p7s", "application/pkcs7-signature")
	mime.AddExtensionType(".p8", "application/pkcs8")
	mime.AddExtensionType(".pak", "application/x-pak")
	mime.AddExtensionType(".par2", "application/x-par2")
	mime.AddExtensionType(".pcap", "application/vnd.tcpdump.pcap")
	mime.AddExtensionType(".pcapng", "application/x-pcapng")
	mime.AddExtensionType(".pcf", "application/x-font-pcf")
	mime.AddExtensionType(".pfa", "application/x-font-type1")
	mime.AddExtensionType(".pfb", "application/x-font-type1")
	mime.AddExtensionType(".pfx", "application/x-pkcs12")
	mime.AddExtensionType(".pgp", "application/pgp-encrypted")
	mime.AddExtensionType(".php3", "application/x-php")
	mime.AddExtensionType(".php4", "application/x-php")
	mime.AddExtensionType(".php5", "application/x-php")
	mime.AddExtensionType(".php", "application/x-php")
	mime.AddExtensionType(".pkcs8", "application/pkcs8")
	mime.AddExtensionType(".pko", "application/ynd.ms-pkipko")
	mime.AddExtensionType(".pkr", "application/pgp-keys")
	mime.AddExtensionType(".pmc", "application/x-perfmon")
	mime.AddExtensionType(".potm", "application/vnd.ms-powerpoint.template.macroEnabled.12")
	mime.AddExtensionType(".potx", "application/vnd.openxmlformats-officedocument.presentationml.template")
	mime.AddExtensionType(".ppam", "application/vnd.ms-powerpoint.addin.macroEnabled.12")
	mime.AddExtensionType(".pps", "application/vnd.ms-powerpoint")
	mime.AddExtensionType(".ppsm", "application/vnd.ms-powerpoint.slideshow.macroEnabled.12")
	mime.AddExtensionType(".ppsx", "application/vnd.openxmlformats-officedocument.presentationml.slideshow")
	mime.AddExtensionType(".ppt", "application/vnd.ms-powerpoint")
	mime.AddExtensionType(".pptm", "application/vnd.ms-powerpoint.presentation.macroEnabled.12")
	mime.AddExtensionType(".pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation")
	mime.AddExtensionType(".prc", "application/x-mobipocket-ebook")
	mime.AddExtensionType(".prf", "application/pics-rules")
	mime.AddExtensionType(".ps", "application/postscript")
	mime.AddExtensionType(".rar", "application/rar")
	mime.AddExtensionType(".rm", "application/vnd.rn-realmedia")
	mime.AddExtensionType(".rng", "application/xml")
	mime.AddExtensionType(".roff", "application/x-troff")
	mime.AddExtensionType(".rpm", "application/x-rpm")
	mime.AddExtensionType(".rtf", "application/rtf")
	mime.AddExtensionType(".rtx", "application/rtf")
	mime.AddExtensionType(".sas", "application/sas")
	mime.AddExtensionType(".scd", "application/x-msschedule")
	mime.AddExtensionType(".sd2", "application/spss")
	mime.AddExtensionType(".sea", "application/x-sea")
	mime.AddExtensionType(".sfs", "application/vnd.squashfs")
	mime.AddExtensionType(".sg", "application/x-sg1000-rom")
	mime.AddExtensionType(".shar", "application/x-shar")
	mime.AddExtensionType(".sig", "application/pgp-signature")
	mime.AddExtensionType(".sis", "application/vnd.symbian.install")
	mime.AddExtensionType(".sit", "application/stuffit")
	mime.AddExtensionType(".skr", "application/pgp-keys")
	mime.AddExtensionType(".smil", "application/smil")
	mime.AddExtensionType(".so", "application/x-sharedlib")
	mime.AddExtensionType(".spl", "application/x-shockwave-flash")
	mime.AddExtensionType(".spo", "application/spss")
	mime.AddExtensionType(".sqsh", "application/vnd.squashfs")
	mime.AddExtensionType(".squashfs", "application/vnd.squashfs")
	mime.AddExtensionType(".sst", "application/vnd.ms-pkicertstore")
	mime.AddExtensionType(".sv4cpio", "application/x-sv4cpio")
	mime.AddExtensionType(".swa", "application/x-director")
	mime.AddExtensionType(".swf", "application/x-shockwave-flash")
	mime.AddExtensionType(".swm", "application/x-ms-wim")
	mime.AddExtensionType(".sxw", "application/vnd.sun.xml.writer")
	mime.AddExtensionType(".t", "application/x-troff")
	mime.AddExtensionType(".tar", "application/tar")
	mime.AddExtensionType(".taz", "application/x-tarz")
	mime.AddExtensionType(".tbz2", "application/x-bzip2")
	mime.AddExtensionType(".tex", "application/x-tex")
	mime.AddExtensionType(".tgz", "application/gzip")
	mime.AddExtensionType(".tlz", "application/x-lzma")
	mime.AddExtensionType(".tnef", "application/ms-tnef")
	mime.AddExtensionType(".toc", "application/x-cdrdao-toc")
	mime.AddExtensionType(".torrent", "application/x-bittorrent")
	mime.AddExtensionType(".tr", "application/x-troff")
	mime.AddExtensionType(".trm", "application/x-msterminal")
	mime.AddExtensionType(".ttc", "application/x-font-ttf")
	mime.AddExtensionType(".ttf", "application/x-font-ttf")
	mime.AddExtensionType(".twb", "application/twb")
	mime.AddExtensionType(".twbx", "application/twb")
	mime.AddExtensionType(".txz", "application/x-xz")
	mime.AddExtensionType(".tzo", "application/x-lzop")
	mime.AddExtensionType(".ufraw", "application/x-ufraw")
	mime.AddExtensionType(".ui", "application/x-designer")
	mime.AddExtensionType(".ustar", "application/x-ustar")
	mime.AddExtensionType(".vbox", "application/x-virtualbox-vbox")
	mime.AddExtensionType(".vbox-extpack", "application/x-virtualbox-vbox-extpack")
	mime.AddExtensionType(".vbs", "application/x-vbscript")
	mime.AddExtensionType(".vdi", "application/x-virtualbox-vdi")
	mime.AddExtensionType(".vhd", "application/x-virtualbox-vhd")
	mime.AddExtensionType(".vmdk", "application/x-virtualbox-vmdk")
	mime.AddExtensionType(".vsd", "application/vnd.visio")
	mime.AddExtensionType(".w3d", "application/x-director")
	mime.AddExtensionType(".wad", "application/x-doom-wad")
	mime.AddExtensionType(".war", "application/x-webarchive")
	mime.AddExtensionType(".wcm", "application/vnd.ms-works")
	mime.AddExtensionType(".wdb", "application/vnd.ms-works")
	mime.AddExtensionType(".wim", "application/x-ms-wim")
	mime.AddExtensionType(".wks", "application/vnd.ms-works")
	mime.AddExtensionType(".wmz", "application/x-ms-wmz")
	mime.AddExtensionType(".woff", "application/font-woff")
	mime.AddExtensionType(".wp4", "application/vnd.wordperfect")
	mime.AddExtensionType(".wp5", "application/vnd.wordperfect")
	mime.AddExtensionType(".wp6", "application/vnd.wordperfect")
	mime.AddExtensionType(".wpd", "application/wordperfect")
	mime.AddExtensionType(".wps", "application/vnd.ms-works")
	mime.AddExtensionType(".wrl", "application/x-mswrite")
	mime.AddExtensionType(".xar", "application/x-xar")
	mime.AddExtensionType(".xla", "application/vnd.ms-excel")
	mime.AddExtensionType(".xlam", "application/vnd.ms-excel.addin.macroEnabled.12")
	mime.AddExtensionType(".xlc", "application/vnd.ms-excel")
	mime.AddExtensionType(".xll", "application/vnd.ms-excel")
	mime.AddExtensionType(".xlm", "application/vnd.ms-excel")
	mime.AddExtensionType(".xls", "application/vnd.ms-excel")
	mime.AddExtensionType(".xlsb", "application/vnd.ms-excel.sheet.binary.macroEnabled.12")
	mime.AddExtensionType(".xlsm", "application/vnd.ms-excel.sheet.macroEnabled.12")
	mime.AddExtensionType(".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	mime.AddExtensionType(".xlt", "application/vnd.ms-excel")
	mime.AddExtensionType(".xltm", "application/vnd.ms-excel.template.macroEnabled.12")
	mime.AddExtensionType(".xltx", "application/vnd.openxmlformats-officedocument.spreadsheetml.template")
	mime.AddExtensionType(".xlw", "application/vnd.ms-excel")
	mime.AddExtensionType(".xps", "application/vnd.ms-xpsdocument")
	mime.AddExtensionType(".xsd", "application/xml")
	mime.AddExtensionType(".xz", "application/x-xz")
	mime.AddExtensionType(".yaml", "application/x-yaml")
	mime.AddExtensionType(".yml", "application/x-yaml")
	mime.AddExtensionType(".z64", "application/x-n64-rom")
	mime.AddExtensionType(".zabw", "application/x-abiword")
	mime.AddExtensionType(".z", "application/x-compress")
	mime.AddExtensionType(".zip", "application/zip")
	mime.AddExtensionType(".zoo", "application/x-zoo")
	mime.AddExtensionType(".zz", "application/zlib")
}

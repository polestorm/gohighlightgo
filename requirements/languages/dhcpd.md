# Wymagania: ISC dhcpd configuration

Źródło: dhcpd.conf(5) manual page (ISC DHCP). Definicja: `highlight/lang/dhcpd.hl`.

| Id | Wymaganie |
|---|---|
| LANG-2931 | Nazwy języka: `dhcpd`, `dhcpdconf`. |
| LANG-2932 | Słowa z listy keyword (`subnet`, `netmask`, `range`, `host`, `option`, `hardware` …) dostają `hl-kw`; słowa z listy builtin (`routers`, `domain-name`, `domain-name-servers`, `broadcast-address`, `true`, `false` …) dostają `hl-builtin`. |
| LANG-2933 | Komentarze `#` do końca wiersza dostają `hl-com`; napisy `"…"` z ucieczkami `hl-esc` dostają `hl-str`. |
| LANG-2934 | Liczby dostają `hl-num`; znaki `;{},:` dostają `hl-op`. |

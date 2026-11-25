# Wymagania: Objective-C

Źródło: Apple „Programming with Objective-C” i norma C (C23) dla części wspólnej. Definicja: `highlight/lang/objectivec.hl`.

| Id | Wymaganie |
|---|---|
| LANG-311 | Nazwy języka: `objectivec`, `objc`, `objective-c`, `m`, `mm`. |
| LANG-312 | Wszystko jak w C (`LANG-012`–`LANG-017`), a dodatkowo słowa `self super in out inout bycopy byref oneway` dostają `hl-kw`, typy `id instancetype BOOL SEL Class IMP NSInteger NSUInteger CGFloat` `hl-type`, `YES NO nil Nil` `hl-builtin`. |
| LANG-313 | Dyrektywy `@interface`, `@implementation`, `@end`, `@property` i inne `@nazwa` dostają `hl-kw`; literał `@"…"` dostaje `hl-str`. |

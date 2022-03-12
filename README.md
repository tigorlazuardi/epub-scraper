# Epub Scraper

Epub Scraper is a tool to scrape websites by it's content and make them into epubs. It uses a domain specific configuration and query selector to search and scrape for those spcific content.

---

1\.  [Disclaimer](#disclaimer)  
2\.  [CHANGELOG](#changelog)  

---

<a name="disclaimer"></a>

## 1\. Disclaimer

This program is given as is. Any problem caused by the user for the usage of this software (be it directly or indirectly, legal issue or non-legal issue) is not the developer's fault.

This program is intended for my personal use. The developer may receive issue notice, but will not promise nor will affirm to fix those issue.

---

<a name="changelog"></a>

## 2\. CHANGELOG

<a name="unreleased"></a>
> [Unreleased]

> Chore
- simplified logging package
- better error message grammar

> Docs
- **changelog:** added changelog
- **readme:** various readme update

> Feat
- **lefthook:** fix markdown-pp not running
- **light-novel-translation:** implemented scraper
- **scraper:** update logger infor
- **scraper:** added generic scraper generator
- **scraper:** added generic scraper
- **scraper-error:** implement display

> Fix
- **light-novel-translation:** fix potential panic on wrong Done() counter


<a name="v0.1.0"></a>
> v0.1.0 - 2022-03-10
> Feat
- **init:** skeleton mappings
- **scraper:** wip for light novel translation
- **scraper:** added marshal json to scrape error
- **scraper:** skeleton update
- **skeleton:** update skeleton
- **unsafe:** added unsafe utilities

> Fix
- **light-novel-translation:** better error message for status code
- **scraper:** fix stack overflow on ScrapeError struct


[Unreleased]: https://github.com/tigorlazuardi/epub-scraper/compare/v0.1.0...HEAD


---
title1: "{{ words 2 " " | capFirst }}"
date: {{ date "2014-11-12" 250 "Jan 2, 06" }}
title2: "{{ words 2 " " | capEach }}"
title3: "{{ words 2 " " | capAll }}"
sents: "{{ sents 3}}"
tags: ["{{words 1 ` ` }}", "{{ words 2 ` ` }}"]
---

# Heading 1
information: {{ paras 8 2}}
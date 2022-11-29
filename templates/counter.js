"use strict";
var ymetrika = {
    {{range .Counters}}
    {{.Name}}:{{.ID}},
};
var metriks = function metriks(label) {
    (function (m, e, t, r, i, k, a) {
        m[i] = m[i] || function () {
            (m[i].a = m[i].a || []).push(arguments);
        };
        m[i].l = 1 * new Date();
        k = e.createElement(t), a = e.getElementsByTagName(t)[0], k.async = 1, k.src = r, a.parentNode.insertBefore(k, a);
    })(window, document, "script", "https://mc.yandex.ru/metrika/tag.js", "ym");
    ym(label, "init", {
        clickmap: true,
        trackLinks: true,
        accurateTrackBounce: true,
        webvisor: true
    });
};
var pathName = document.location.pathname;
var pathNameArr = pathName.split('/');
var version = pathNameArr.filter(function (item) {
    if (item in ymetrika) return item;
}).join();
if (version in ymetrika) {
    metriks(ymetrika[version]);
};
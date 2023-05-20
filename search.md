---
title : "Search"
---


# Project Search

<noscript>JavaScript must be enabled for search to work</noscript>

<div id="searchbox">

<input id="query" type="text" value="" size="60" placeholder="search blog"> <input id="search" type="submit" value="Search"></div>

<hr>

<div id="results"></div>


<link href="./pagefind/pagefind-ui.css" rel="stylesheet">

<script src="./pagefind/pagefind-ui.js" type="text/javascript"></script>

<script>
    window.addEventListener('DOMContentLoaded', (event) => {
        new PagefindUI({ 
            element: "#search",
            baseUrl: "./"
        });
    });
</script>




# :mag: checkdns.run

checkdns.run is an extremely simple DNS lookup service, aiming to allow people to quickly
check DNS records on various nameservers.

Currently supports queries on:
* A
* CNAME
* TXT
* SOA

You can use the UI to look up a DNS record, or you can go directly to results using URL params. 
For example, to look up the A record for `example.com` using Google's DNS at `8.8.8.8`, you can go directly
to this URL:

`https://checkdns.run/A/example.com@8.8.8.8`

## Tech

* Web server built with Go, using [miekg/dns](https://github.com/miekg/dns) for DNS queries 
and [gorilla/mux](https://github.com/gorilla/mux) for routing
* CSS with [tailwind](https://tailwindcss.com/)
* Some small bits of interactivity with [alpine.js](https://alpinejs.dev/)
* [fly.io](https://fly.io/) for hosting and cert management
* [Caddy](https://caddyserver.com/) for static files and general webserver loveliness

## TODO

* Add the rest of the records that miekg/dns supports (basically, all of them)
* Add ability to check propogation of records
* Add permalink to results pages



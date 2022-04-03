gpxaltitude
===========

Gpxaltitude is a simple tool which allows to fix altitudes in GPS traces encoded in [GPS Exchange Format](https://www.topografix.com/gpx.asp) (GPX).

Usage
-----

```
gpxaltitude [ -v ] [ -a n | -r n ] file\n
```

 - The `-a` (absolute) parameter sets the starting altitude to a known absolute altitude and fixes all the following altitudes accordingly.
 - The `-r` (relative) parameter applies a positive or negative relative change to the altitudes.
 - The `-v` parameter increases verbosity on stderr.

When no parameter is given, altitudes are simply rounded to the nearest centimeter.

Examples
--------

Decrease the altitudes values by 24 meters:

```
gpxaltitude -d -24 trace.gpx >trace_fixed.gpx
```

Set the starting altitude to 1257 meters:

```
gpxaltitude -a 1257 trace.gpx >trace_fixed.gpx
```

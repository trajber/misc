#!/usr/bin/env python

import os
import re
import sys

def exit_with_error():
    print "Please fix it... bye!"
    sys.exit(1)

# you can change these extensions if you want
movies = filter(lambda x: x.endswith('.avi'), os.listdir('.'))
subtitles = filter(lambda x: x.endswith('.srt'), os.listdir('.'))

movies_and_subtitles = []

for movie_name in movies:
    # the following pattern is the mostly used: S<somedigit>E<somedigit>
    match_movie = re.search("S\d+E\d+", movie_name,re.IGNORECASE)

    thesubtitle = filter(lambda x: re.search(match_movie.group(), x, re.IGNORECASE), subtitles)

    if (len(thesubtitle) == 0):
        print "Subtitle not found for file [%s]" % movie_name
        exit_with_error()

    if (len(thesubtitle) > 1):
        print "More than one subtitle for file [%s]" % movie_name
        print "Candidates are:" + str(thesubtitle)
        exit_with_error()

    movies_and_subtitles.append((movie_name,thesubtitle[0]))

for ms in movies_and_subtitles:
    print "Renaming [%s] to [%s]" % (ms[0][:-4], ms[1][:-4])
    os.rename(ms[1], ms[0][:-4] + ".srt")

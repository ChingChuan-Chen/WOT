library(pipeR)
library(DBI)
library(ROracle)
library(data.table)
source("data_prepare.R", local = TRUE)

getOraConn <- function(){
  host <- "192.168.1.113"
  port <- 1539
  sid <- "orcl"
  connectString <- sprintf("(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=%s)(PORT=%i))(CONNECT_DATA=(SID=%s)))",
                           host, port, sid)
  dbConnect(dbDriver("Oracle"), username = "webdbtool",
            password = "f15cim2w3e4r", dbname = connectString)
}


# name.basics.tsv.gz – Contains the following information for names:
#   nconst (string) - alphanumeric unique identifier of the name/person
#   primaryName (string)– name by which the person is most often credited
#   birthYear – in YYYY format
#   deathYear – in YYYY format if applicable, else ‘\N’
#   primaryProfession (array of strings)– the top-3 professions of the person
#   knownForTitles (array of tconsts) – titles the person is known for

# title.akas.tsv.gz - Contains the following information for titles:
#   titleId (string) - a tconst, an alphanumeric unique identifier of the title
#   ordering (integer) – a number to uniquely identify rows for a given titleId
#   title (string) – the localized title
#   region (string) - the region for this version of the title
#   language (string) - the language of the title
#   types (array) - Enumerated set of attributes for this alternative title. One or more of the following: "alternative", "dvd", "festival", "tv", "video", "working", "original", "imdbDisplay". New values may be added in the future without warning
#   attributes (array) - Additional terms to describe this alternative title, not enumerated
#   isOriginalTitle (boolean) – 0: not original title; 1: original title

# title.basics.tsv.gz - Contains the following information for titles:
#   tconst (string) - alphanumeric unique identifier of the title
#   titleType (string) – the type/format of the title (e.g. movie, short, tvseries, tvepisode, video, etc)
#   primaryTitle (string) – the more popular title / the title used by the filmmakers on promotional materials at the point of release
#   originalTitle (string) - original title, in the original language
#   isAdult (boolean) - 0: non-adult title; 1: adult title
#   startYear (YYYY) – represents the release year of a title. In the case of TV Series, it is the series start year
#   endYear (YYYY) – TV Series end year. ‘\N’ for all other title types
#   runtimeMinutes – primary runtime of the title, in minutes
#   genres (string array) – includes up to three genres associated with the title

# title.crew.tsv.gz – Contains the director and writer information for all the titles in IMDb. Fields include:
#   tconst (string) - alphanumeric unique identifier of the title
#   directors (array of nconsts) - director(s) of the given title
#   writers (array of nconsts) – writer(s) of the given title

# title.episode.tsv.gz – Contains the tv episode information. Fields include:
#   tconst (string) - alphanumeric identifier of episode
#   parentTconst (string) - alphanumeric identifier of the parent TV Series
#   seasonNumber (integer) – season number the episode belongs to
#   episodeNumber (integer) – episode number of the tconst in the TV series

# title.principals.tsv.gz – Contains the principal cast/crew for titles
#   tconst (string) - alphanumeric unique identifier of the title
#   principalCast (array of nconsts) – title’s top-billed cast/crew

# title.ratings.tsv.gz – Contains the IMDb rating and votes information for titles
#   tconst (string) - alphanumeric unique identifier of the title
#   averageRating – weighted average of all the individual user ratings
#   numVotes - number of votes the title has received

datList <- list.files("tsv_gz_data", "\\.tsv$", full.names = TRUE) %>>%
  lapply(fread, sep = "\t", na.string = c("NA", "\\N"), quote = "")



oraConn <- getOraConn()


dbDisconnect(oraConn)

library(pipeR)
library(httr)
library(xml2)
library(R.utils)

##### get urls of data #####
dataUrls <- GET("https://datasets.imdbws.com/") %>>%
  content(encoding = "UTF-8") %>>% xml_find_all("//ul/a") %>>% xml_attr("href")
stopifnot(length(dataUrls) == 7L)

if (!dir.exists("tsv_gz_data"))
  dir.create("tsv_gz_data")

##### download data #####
for (i in seq_along(dataUrls)) {
  fn <- file.path("tsv_gz_data", basename(dataUrls[i]))
  if (!file.exists(fn))
    download.file(dataUrls[i], fn, quiet = TRUE)
}

##### decompress data #####
list.files("tsv_gz_data/", full.names = TRUE, pattern = "\\.gz$") %>>%
  sapply(function(gzfn){
    destname <- substring(gzfn, 1L, nchar(gzfn) - 3)
    if (!file.exists(destname))
      decompressFile(gzfn, destname, "gz", gzfile, remove = FALSE)
  }) %>>% invisible

##### exit script #####
rm(dataUrls, i, fn)

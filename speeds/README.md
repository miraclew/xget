Downloader procedure:

Main
1. Start with initial file URL
2. Query resource with URL
3. If resource exist, goto GetResource
4. Otherwise, goto NoResource

NoResource
1. Download with original URL
2. If source is http and range is supported, use multi-connection
3. Report Resource, Source, Chunks to server

GetResource
1. Query sources from api
2. Get chunks from api
3. Start download chunks
4. Merge chunks


Client Design

1. use sqlite3 to store downloading data
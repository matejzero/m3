	options m3db.FileWriterOptions,
		options = NewFileWriterOptions()
	shardDir := ShardDirPath(w.filePathPrefix, shard)
	w.checkpointFilePath = filesetPathFromTime(shardDir, blockStart, checkpointFileSuffix)
			filesetPathFromTime(shardDir, blockStart, infoFileSuffix):   &infoFd,
			filesetPathFromTime(shardDir, blockStart, indexFileSuffix):  &indexFd,
			filesetPathFromTime(shardDir, blockStart, dataFileSuffix):   &dataFd,
			filesetPathFromTime(shardDir, blockStart, digestFileSuffix): &digestFd,
	return OpenWritable(filePath, w.newFileMode)
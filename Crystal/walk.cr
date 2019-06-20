def main(dirpath)
  count = 0
  Dir[dirpath + "/**/*"].each do |path|
      filetype = File.info(path, follow_symlinks: false).type
      if filetype.file? && path.downcase.ends_with?(".jpg")
        count += 1
      end
  end
  count
end

puts main(ARGV[0])

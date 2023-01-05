import glob
import json

path = 'db/**/*.json'
for file in glob.glob(path, recursive=True):
  vals = json.load(open(file))
  json.dump(vals, open(file, 'w'), indent=2)
  if len(vals) != len(list(set(vals))):
    print(f'Duplicates found in file: {file}')
    vals_no_dups = list(dict.fromkeys(vals))
    print(f'Original length: {len(vals)}')
    print(f'New length: {len(vals_no_dups)}')
    json.dump(vals_no_dups, open(file, 'w'), indent=2)

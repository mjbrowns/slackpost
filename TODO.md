# TO DO

* Clean up error handling
   * Errors and logging should go to stdout
   * Convert all actions to use RunE with consistent return handling
* Create new pipe mode -p | --pipe
   * init will send the json object to stdout
   * send will read json object from stdin
   * others will read json from stdin, process and send output to stdout
* add --dry-run function that will display a formatted json output but not actually process (send)

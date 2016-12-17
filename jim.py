# A simple multi-domain static web server based on the FileServer
# in the golang standard net/http library.
#
# Usage:
#   python ../rye/rye.py build jim.py       # to compile.
#   ./jim.bin localhost:8080 webs.conf      # to run.
#
#      The first arg is the host:port to bind to.
#
#      The second arg is the config file, which has lines of the form
#          webDir pattern1 pattern2 pattern3 ...
#      where pattern syntax is defined by
#        https://golang.org/pkg/net/http/#ServeMux
#      and webDir is the root of the dir tree to be served.
#
# jim.py is Rye code.  Compile with github.com/strickyak/rye/.
#
# Or you may use pre-transpiled go code like this:
#   go run rye__/jim/jim/ryemain.go  :8080 webs.conf

from go import net/http, os, time

def main(args):
  bindHostColonPort, confFile = args
  print >>os.Stderr, '# Jim the Server'

  for line in open(confFile).read().split('\n'):
    command = line.split('#')[0]  # Ignore comments after '#'.
    words = command.split()

    if words:
      webDir = words[0]
      for pattern in words[1:]:
        def makeLoggingWrapper(pattern, webDir, handler):
          def loggingWrapper(w, r):
            print >>os.Stderr, '%s (%s == %s) %q' % (
                time.Now().Format(time.Stamp), webDir, pattern, r.RequestURI)
            handler.ServeHTTP(w, r)
          return loggingWrapper

        handler = http.FileServer(go_cast(http.Dir, webDir))
        http.HandleFunc(pattern, makeLoggingWrapper(pattern, webDir, handler))
        print >>os.Stderr, '# %q handles %q' % (webDir, pattern)

  print >>os.Stderr, '# Starting Server at %s' % time.Now().Format(time.Stamp)
  http.ListenAndServe(bindHostColonPort, None)

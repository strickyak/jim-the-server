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

from go import net/http, os, path/filepath, syscall, time

# TODO: See https://play.golang.org/p/dXBizm4xl3
def CheckPublicPermissions(p):
  p = filepath.Clean(p)
  if (os.Stat(p).Mode() & 0004) != 0004:
    raise 'File not public readable: 0%o %q' % (os.Stat(p).Mode(), p)
  d = filepath.Dir(p)
  while d != '.' and d != '/' and d != '':
    if (os.Stat(d).Mode() & 0005) != 0005:
      raise 'Dir not public readable & executable: 0%o %q' % (os.Stat(d).Mode(), d)
    d = filepath.Dir(d)

def main(args):
  bindHostColonPort, confFile = args
  print >>os.Stderr, '# Jim the Server'

  for line in open(confFile).read().split('\n'):
    command = line.split('#')[0]  # Ignore comments after '#'.
    words = command.split()

    if words:
      webDir = words[0]
      for pattern in words[1:]:
        def makeWrapper(pattern, webDir, handler):
          def wrapper(w, r):
            print >>os.Stderr, '%s (%s == %s) %q' % (
                time.Now().Format(time.Stamp), webDir, pattern, r.RequestURI)
            try:
              CheckPublicPermissions(filepath.Join(webDir, r.URL.Path))
            except as e:
              http.NotFound(w, r)
              print >> os.Stderr, 'ERROR %q (%s == %s) %q' % (e, webDir, pattern, r.RequestURI)
              return
            handler.ServeHTTP(w, r)
          return wrapper

        handler = http.FileServer(go_cast(http.Dir, webDir))
        http.HandleFunc(pattern, makeWrapper(pattern, webDir, handler))
        print >>os.Stderr, '# %q handles %q' % (webDir, pattern)

  print >>os.Stderr, '# Starting Server at %s' % time.Now().Format(time.Stamp)

  http.ListenAndServe(bindHostColonPort, None)

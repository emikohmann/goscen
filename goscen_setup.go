package goscen

func init() {
    scoring := read()

    scoring.check()

    scoring.load()

    scoring.success()

    scoring.serve()
}

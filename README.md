# Lugbúrz

Lugbúrz executes Ansible jobs remotely.  The Ansible runtime is distributed
and executed through Docker images, which keeps dependencies to a minimum.

The operator will interact through a kubectl-like CLI.

> Barad-dûr, the “Dark Tower,” is a fictional place in J. R. R. Tolkien's
  Middle-earth writings and is described in The Lord of the Rings, The Silmarillion,
  and other works. It is an enormous fortress of the Dark Lord Sauron, whence he
  rules the volcanic and barren land of Mordor. Located in northwest Mordor, near
  Mount Doom, the Eye of Sauron keeps watch over Middle-earth from its highest tower.
  The name is pronounced "Ba'rad doorr" with emphasis placed on the "rr."  The
  Lieutenant of Barad-dûr is the Mouth of Sauron, who acts as an ambassador and
  herald for Mordor and Sauron.  Barad-dûr was called "Lugbúrz" in the Black Speech
  of Mordor, which also translates as "Dark Tower"; it is composed of Lug = Tower
  and Búrz = Dark.

https://en.wikipedia.org/wiki/Barad-dûr

## Dependencies

    $ go get github.com/golang/dep/cmd/dep
    $ go get github.com/jteeuwen/go-bindata
    $ go get github.com/mitchellh/gox

## Building

    $ make build
    $ tree .build/

## Testing

    $ make test

## License

MIT

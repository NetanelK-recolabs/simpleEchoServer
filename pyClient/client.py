#!/usr/bin/env python3

import asyncio
from grpclib.client import Channel
import gen.proto.python.echoapis.echo.v1 as v1


HOST = '127.0.0.1'  # The server's hostname or IP address
PORT = 65432  # The port used by the server

'''Simple python client'''


async def main() -> None:
    channel = Channel(host=HOST, port=PORT)
    service = v1.EchoServiceStub(channel)

    response = await service.echo(msg="Nate")
    print(f"echo {response.msg=}")
    response = await service.echo_many(msg="Yada", repeat=3)
    print(f"echo many {response.msg=}")
    # don't forget to close the channel when done!
    channel.close()


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())

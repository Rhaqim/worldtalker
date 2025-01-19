import asyncio
import signal
import time
from concurrent import futures

import grpc

from wtai.proto import translate_pb2, translate_pb2_grpc
from wtai.translator.translate import AITranslator


# Create a class to define the server functions, derived from translate_pb2_grpc.TranslatorServiceServicer
class TranslatorServiceServicer(translate_pb2_grpc.TranslatorServicer):
    async def Translate(self, request, context):
        """
        Handles the Translate RPC request.

        Args:
            request (TranslateRequest): The request containing the message and language details.
            context (grpc.aio.ServicerContext): The context of the RPC.

        Returns:
            TranslateResponse: The response containing the translated message and metadata.
        """
        translator = AITranslator()

        try:
            # Perform translation (await if translate is an async method)
            translated_message = await translator.translate(
                request.message, request.language_source, request.language_target
            )

            # Create a response
            return translate_pb2.TranslateResponse(
                translated_message=translated_message,
                metadata=f"Translated from {request.language_source} to {request.language_target}",
            )

        except Exception as e:
            # Handle errors and return an appropriate response
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return translate_pb2.TranslateResponse(
                translated_message="", metadata="An error occurred during translation."
            )


async def serve():
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))

    translate_pb2_grpc.add_TranslatorServicer_to_server(
        TranslatorServiceServicer(), server
    )

    # server.add_insecure_port("[::]:50051")
    server.add_insecure_port("0.0.0.0:50051")
    await server.start()

    print("Server started on port 50051")

    # Gracefully wait for termination
    stop_event = asyncio.Event()

    def handle_signal():
        print("\nGracefully shutting down...")
        stop_event.set()

    # Register signal handlers for termination
    loop = asyncio.get_running_loop()
    loop.add_signal_handler(signal.SIGINT, handle_signal)
    loop.add_signal_handler(signal.SIGTERM, handle_signal)

    try:
        await stop_event.wait()
    finally:
        await server.stop(grace=5)  # Gracefully stop the server, allowing pending RPCs to finish
        print("Server shut down.")

    await server.wait_for_termination()


if __name__ == "__main__":
    asyncio.run(serve())

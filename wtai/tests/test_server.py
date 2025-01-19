from unittest.mock import AsyncMock, patch

import grpc
import pytest
from grpc import aio

from wtai.proto import translate_pb2, translate_pb2_grpc
from wtai.server import TranslatorServiceServicer


@pytest.mark.asyncio
async def test_translate_success():
    # Arrange
    request = translate_pb2.TranslateRequest(message="Hello", language_source="en", language_target="es")
    context = AsyncMock()

    with patch('wtai.server.AITranslator') as mock_translator:
        mock_translator.return_value.translate = AsyncMock(return_value="Hola")

        servicer = TranslatorServiceServicer()

        # Act
        response = await servicer.Translate(request, context)

        # Assert
        assert response.translated_message == "Hola"
        assert response.metadata == "Translated from en to es"
        mock_translator.return_value.translate.assert_called_once_with("Hello", "en", "es")
        context.set_code.assert_not_called()
        context.set_details.assert_not_called()


@pytest.mark.asyncio
async def test_translate_error():
    # Arrange
    request = translate_pb2.TranslateRequest(message="Hello", language_source="en", language_target="es")
    context = AsyncMock()

    with patch('wtai.server.AITranslator') as mock_translator:
        mock_translator.return_value.translate = AsyncMock(side_effect=Exception("Translation error"))

        servicer = TranslatorServiceServicer()

        # Act
        response = await servicer.Translate(request, context)

        # Assert
        assert response.translated_message == ""
        assert response.metadata == "An error occurred during translation."
        mock_translator.return_value.translate.assert_called_once_with("Hello", "en", "es")
        context.set_code.assert_called_once_with(grpc.StatusCode.INTERNAL)
        context.set_details.assert_called_once_with("Translation error")


@pytest.mark.asyncio
@pytest.mark.parametrize(
    "message, source_lang, target_lang, translated_message, metadata",
    [
        ("Hello", "en", "es", "Hola", "Translated from en to es"),
        ("Goodbye", "en", "fr", "Au revoir", "Translated from en to fr"),
        ("Thank you", "en", "de", "Danke", "Translated from en to de"),
    ],
)
async def test_translate_parametrized(message, source_lang, target_lang, translated_message, metadata):
    # Arrange
    request = translate_pb2.TranslateRequest(message=message, language_source=source_lang, language_target=target_lang)
    context = AsyncMock()

    with patch('wtai.server.AITranslator') as mock_translator:
        mock_translator.return_value.translate = AsyncMock(return_value=translated_message)

        servicer = TranslatorServiceServicer()

        # Act
        response = await servicer.Translate(request, context)

        # Assert
        assert response.translated_message == translated_message
        assert response.metadata == metadata
        mock_translator.return_value.translate.assert_called_once_with(message, source_lang, target_lang)
        context.set_code.assert_not_called()
        context.set_details.assert_not_called()
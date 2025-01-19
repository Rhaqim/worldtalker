from unittest.mock import AsyncMock, patch

import pytest

from wtai.translator.translate import AITranslator


@pytest.mark.asyncio
async def test_translate_success():
    with patch('wtai.translator.translate.Translator') as mock_translator:
        # Use AsyncMock for async methods
        mock_translator.return_value.translate = AsyncMock(return_value=AsyncMock(text="Hola"))
        
        translator = AITranslator()
        result = await translator.translate("Hello", "en", "es")
        
        assert result == "Hola"
        mock_translator.return_value.translate.assert_called_once_with("Hello", src="en", dest="es")

@pytest.mark.asyncio
async def test_translate_error():
    with patch('wtai.translator.translate.Translator') as mock_translator:
        # Simulate an exception for the async method
        mock_translator.return_value.translate = AsyncMock(side_effect=Exception("Translation error"))
        
        translator = AITranslator()
        result = await translator.translate("Hello", "en", "es")
        
        assert result == "Error: Translation error"
        mock_translator.return_value.translate.assert_called_once_with("Hello", src="en", dest="es")
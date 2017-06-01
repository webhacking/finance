class InvalidArgumentException(Exception):
    def __init__(self, arg):
        message = 'Invalid argument: {}'.format(arg)
        super(InvalidArgumentException, self).__init__(message)


class AccountNotFoundException(Exception):
    def __init__(self, arg):
        message = 'Non-existing account: {}'.format(arg)
        super(self.__class__, self).__init__(message)


class AssetNotFoundException(Exception):
    def __init__(self, arg):
        message = 'Non-existing asset: {}'.format(arg)
        super(self.__class__, self).__init__(message)


class AssetValueUnavailableException(Exception):
    pass


class InvalidTargetAssetException(Exception):
    pass
